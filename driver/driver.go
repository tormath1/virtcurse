package driver

// Driver interface represents the
// middleware between the APIs and the TUI
type Driver interface {
	// Close close the current driver connection
	Close() error

	// ListDomain returns a list of domains using
	// options
	ListDomain(opts Options) ([]Domain, error)

	// GetDomain returns a domain from its name
	GetDomain(name string) (Domain, error)
}
