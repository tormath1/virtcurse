package driver

type ListFlag int

const (
	All ListFlag = iota
	Shutdown
	Active
)

// Options is the helper to handle
// driver options for requests
type Options struct {
	ListOptions ListFlag
}
