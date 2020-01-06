package maimai

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/levigross/grequests"
)

const (
	DefaultUA       = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"
	DefaultAreaCode = "+86"
	loginURL        = "https://acc.maimai.cn/login"
)

var (
	msession    *grequests.Session
	accessToken string //
	uid         string //maimai unique userid
)

func init() {
	msession = grequests.NewSession(nil)
}

//login to maimai.cn and get access token
//POST https://acc.maimai.cn/login
// pa=+86	area code
// m=1xxxxxxxxxx	phone number
// p=xxxxx		password
//once succeed, it will set uid and accessToken correctly, which are necessary for getting gossips
func Login(phonenumber string, password string) error {
	var resp *http.Response
	var err error
	v := url.Values{}
	v.Add("pa", DefaultAreaCode)
	v.Add("m", phonenumber)
	v.Add("p", password)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err = client.PostForm(loginURL, v)
	if err != nil {
		return fmt.Errorf("failed to post login form, %v", err)
	}

	//usually returns with 302 and go to /set_auth_cookie
	if resp.StatusCode != http.StatusFound {
		return fmt.Errorf("unexpected response code, code=%d", resp.StatusCode)
	}

	//should be https://maimai.cn/set_auth_cookie?...
	newlocation := resp.Header.Get("location")
	resp, err = client.Get(newlocation)
	if err != nil {
		return fmt.Errorf("failed to get access token, %v", err)
	}

	//get token/uid from response
	for _, cookie := range resp.Cookies() {
		switch cookie.Name {
		case "u":
			uid = cookie.Value
		case "access_token":
			accessToken = cookie.Value
		}
	}

	if len(uid) == 0 || len(accessToken) == 0 {
		return fmt.Errorf("invalid uid or access token,uid=%s, token=%s", uid, accessToken)
	}

	return nil
}
