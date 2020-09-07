package libvirt

import (
	libvirt "github.com/libvirt/libvirt-go"

	"github.com/pkg/errors"
	"github.com/tormath1/virtcurse/driver"
)

type Libvirt struct{ conn *libvirt.Connect }

// NewLibvirt return a driver implementing
// the driver interface
func NewLibvirt(URI string) (driver.Driver, error) {
	c, err := libvirt.NewConnect(URI)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create libvirt client")
	}
	return &Libvirt{conn: c}, nil
}

func (l Libvirt) ListDomain(opts driver.Options) ([]driver.Domain, error) {
	var (
		listFlags libvirt.ConnectListAllDomainsFlags
		domains   []driver.Domain
	)
	switch opts.ListOptions {
	case driver.All:
		listFlags = libvirt.CONNECT_LIST_DOMAINS_ACTIVE +
			libvirt.CONNECT_LIST_DOMAINS_INACTIVE +
			libvirt.CONNECT_LIST_DOMAINS_SHUTOFF +
			libvirt.CONNECT_LIST_DOMAINS_PAUSED
	case driver.Shutdown:
		listFlags = libvirt.CONNECT_LIST_DOMAINS_SHUTOFF
	case driver.Active:
		listFlags = libvirt.CONNECT_LIST_DOMAINS_ACTIVE
	}
	doms, err := l.conn.ListAllDomains(listFlags)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list all domains")
	}
	domains = make([]driver.Domain, 0)
	for _, dom := range doms {
		n, err := dom.GetName()
		if err != nil {
			// TODO: Should we skip this error and use a default name
			return nil, errors.Wrap(err, "unable to get domain's name")
		}
		domains = append(domains, Domain{name: n})
	}
	return domains, nil
}

func (l Libvirt) Close() error {
	if _, err := l.conn.Close(); err != nil {
		return errors.Wrap(err, "unable to close connection correctly")
	}
	return nil
}
