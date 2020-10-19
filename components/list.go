package components

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/tormath1/virtcurse/driver"
)

// ListDomains is the component displaying
// the running domains
type ListDomains struct {
	*widgets.List
	d driver.Driver
}

// NewListDomains initializes the ListDomains component
func NewListDomains(d driver.Driver) *ListDomains {
	ld := &ListDomains{
		List: widgets.NewList(),
		d:    d,
	}

	domains, _ := ld.d.ListDomain(driver.Options{})
	rows := make([]string, len(domains))

	for _, domain := range domains {
		name, _ := domain.Name()
		rows = append(rows, name)
	}

	width, height := ui.TerminalDimensions()

	ld.Title = "Running domains"
	ld.WrapText = false
	ld.Rows = rows
	ld.SetRect(0, 0, width/2, height)
	return ld
}
