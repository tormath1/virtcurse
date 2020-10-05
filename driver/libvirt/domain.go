package libvirt

import (
	libvirt "github.com/libvirt/libvirt-go"
	"github.com/pkg/errors"
)

type Domain struct {
	dom libvirt.Domain
}

func (d Domain) Name() (string, error) {
	name, err := d.dom.GetName()
	if err != nil {
		// TODO: Should we skip this error and use a default name
		return "", errors.Wrap(err, "unable to get domain's name")
	}
	return name, nil
}

func (d Domain) XML() (string, error) {
	xml, err := d.dom.GetXMLDesc(0)
	if err != nil {
		return "", errors.Wrap(err, "unable to get xml description")
	}
	return xml, nil
}
