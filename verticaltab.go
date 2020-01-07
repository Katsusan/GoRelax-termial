package main

import (
	"image"
	"sync"

	ui "github.com/gizak/termui/v3"
)

/*
type Drawable interface {
	GetRect() image.Rectangle
	SetRect(int, int, int, int)
	Draw(*Buffer)
	sync.Locker
}
*/

//https://www.key-shortcut.com/en/writing-systems/35-symbols/arrows
const (
	DOWNWARD_ARROW = "⮩"
)

type VerticalTabPane struct {
	ui.Block
	TableNames       []string
	ActiveTabIndex   int
	ActiveTabStyle   ui.Style
	InactiveTabStyle ui.Style
	sync.Mutex
}

func NewVerticalTabPane(names ...string) *VerticalTabPane {
	return &VerticalTabPane{
		Block:            *ui.NewBlock(),
		TableNames:       names,
		ActiveTabStyle:   ui.Theme.Tab.Active,
		InactiveTabStyle: ui.Theme.Tab.Inactive,
	}
}

func (tab *VerticalTabPane) FocusUp() {
	if tab.ActiveTabIndex > 0 {
		tab.ActiveTabIndex--
	}
}

func (tab *VerticalTabPane) FocusDown() {
	if tab.ActiveTabIndex < len(tab.TableNames)-1 {
		tab.ActiveTabIndex++
	}
}

func (tab *VerticalTabPane) Draw(buf *ui.Buffer) {
	tab.Block.Draw(buf)

	yCoordinate := tab.Inner.Min.Y
	for i, name := range tab.TableNames {
		styles := tab.InactiveTabStyle
		if i == tab.ActiveTabIndex {
			styles = tab.ActiveTabStyle
		}
		buf.SetString(
			ui.TrimString(name, tab.Inner.Max.X-tab.Inner.Min.X),
			styles,
			image.Pt(yCoordinate, tab.Inner.Min.X),
		)

		yCoordinate++

		if i < len(tab.TableNames) && yCoordinate < tab.Inner.Max.Y {
			buf.SetCell(
				ui.NewCell(ui.HORIZONTAL_LINE, ui.NewStyle(ui.ColorYellow)),
				image.Pt(tab.Inner.Min.X, yCoordinate),
			)
		}

		yCoordinate++
	}
}

func initTab() {

}
