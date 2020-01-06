package maimai

import (
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	m := os.Getenv("MAIMAI_MOB") //get your mobile phone number
	p := os.Getenv("MAIMAI_PWD") //get your password

	if len(m) == 0 || len(p) == 0 {
		t.Fatal("MAIMAI_MOB and MAIMAI_PWD must be set")
	}

	if err := Login(m, p); err != nil {
		t.Fatalf("login failed, error=%v", err)
	}
}
