package components

import (
	ui "github.com/gizak/termui/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/tormath1/virtcurse/driver/libvirt"
)

// Executes the "root" component
func Execute(cmd *cobra.Command) error {
	libvirtURI, err := cmd.Flags().GetString("libvirt-uri")
	if err != nil {
		return errors.Wrap(err, "unable to get libvirt-uri")
	}

	d, err := libvirt.NewLibvirt(libvirtURI)
	if err != nil {
		return errors.Wrap(err, "unable to init libvirt client")
	}

	if err := ui.Init(); err != nil {
		return errors.Wrap(err, "unable to init UI")
	}
	defer ui.Close()

	components := []ui.Drawable{
		NewListDomains(d),
		NewGetConfig(d),
	}

	for _, component := range components {
		ui.Render(component)
	}

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}

	return nil
}
