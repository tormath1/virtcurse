package libvirt

type Domain struct{ name string }

func (d Domain) Name() string {
	return d.name
}
