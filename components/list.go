package components

import (
	"github.com/rivo/tview"
)

// NewListWidget initializes the List component
func NewListWidget(title string, items []string) *tview.List {
	list := tview.NewList()
	list.ShowSecondaryText(false).
		SetTitle(title).
		SetBorder(true)

	for _, item := range items {
		list.AddItem(item, "", 0, nil)
	}

	return list
}
