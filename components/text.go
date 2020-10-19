package components

import (
	"github.com/rivo/tview"
)

// NewTextWidget initializes the Text component
func NewTextWidget(title, content string) *tview.TextView {
	txt := tview.NewTextView()
	txt.SetScrollable(true).
		SetWrap(true).
		SetText(content).
		SetTitle(title).
		SetBorder(true)

	return txt
}
