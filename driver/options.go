package driver

// ListFlag is an optional flag
// to pass to the ListDomain method
type ListFlag int

const (
	// All lists all the domains
	All ListFlag = iota
	// Shutdown lists only shutdowned domains
	Shutdown
	// Active lists only active domains
	Active
)

// Options is the helper to handle
// driver options for requests
type Options struct {
	// ListOptions is the options to pass
	// to the ListDomain methods
	ListOptions ListFlag
}
