package components

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/tormath1/virtcurse/driver"
)

// GetConfig is the config component
// used to display the XML config of a domain
type GetConfig struct {
	*widgets.Paragraph
	d driver.Driver
}

// NewGetConfig initializes a GetConfig component
func NewGetConfig(d driver.Driver) *GetConfig {
	gc := &GetConfig{
		Paragraph: widgets.NewParagraph(),
		d:         d,
	}

	// Fake the call to one particular domain
	// just for testing
	doms, _ := d.ListDomain(driver.Options{})
	var dom driver.Domain
	if len(doms) > 0 {
		dom = doms[0]
	} else {
		return nil
	}

	xml, _ := dom.XML()
	gc.Text = xml

	width, height := ui.TerminalDimensions()
	gc.SetRect(width/2, 0, width, height)

	return gc
}
