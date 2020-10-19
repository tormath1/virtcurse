package components

import (
	"fmt"

	"github.com/rivo/tview"

	"github.com/tormath1/virtcurse/driver"
	"github.com/tormath1/virtcurse/driver/libvirt"
)

// Execute the "root" component
func Execute(libvirtURI string) error {
	d, err := libvirt.NewLibvirt(libvirtURI)
	if err != nil {
		return fmt.Errorf("unable to init libvirt client: %w", err)
	}
	app := tview.NewApplication()
	flex := tview.NewFlex()

	domains, err := listDomains(d)
	if err != nil {
		return fmt.Errorf("unable to list domains: %w", err)
	}

	list := NewListWidget("Running Machines", domains)
	list.SetSelectedFunc(func(idx int, txt, sec string, shc rune) {
		dom, _ := d.GetDomain(txt)
		xml, _ := dom.XML()
		flex.AddItem(NewTextWidget("Machine Configuration", xml), 0, 1, true)
	})
	flex.AddItem(list, 0, 1, true)

	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		return fmt.Errorf("unable to create application: %w", err)
	}

	return nil
}

func listDomains(d driver.Driver) ([]string, error) {
	domains, err := d.ListDomain(driver.Options{})
	if err != nil {
		return nil, fmt.Errorf("unable to create application: %w", err)
	}
	doms := make([]string, len(domains))
	for i, domain := range domains {
		// TODO: log the error
		if name, err := domain.Name(); err == nil {
			doms[i] = name
		} else {
			doms[i] = "n/a"
		}
	}
	return doms, nil
}
