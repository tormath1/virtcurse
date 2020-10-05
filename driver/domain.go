package driver

// Domain is the common interface
// representing domains
type Domain interface {
	// Name is the name of the domain
	Name() (string, error)

	// XML returns the XML description of the domain
	XML() (string, error)
}
