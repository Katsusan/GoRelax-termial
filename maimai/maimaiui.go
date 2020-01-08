package maimai

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type MaiUI struct {
	Tabs    widgets.TabPane
	Gossips widgets.List //职言区
	Feeds   widgets.List //实名区
}

func NewMaiUI() {
	var maiui MaiUI
	maiui.Tabs.TabNames = []string{"职言", "实名"}

	gossipUI := widgets.NewList()
	gossipUI.Title = "职言"
	gossipUI.TextStyle = ui.NewStyle(ui.ColorYellow)

	FeedUI := widgets.NewList()
	FeedUI.Title = "实名"
	FeedUI.TextStyle = ui.NewStyle(ui.ColorYellow)

}
