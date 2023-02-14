package ipcsdk

// Status defines the different states in which a subnet can be.
type Status int64

const (
	Instantiated Status = iota
	Active
	Inactive
	Terminating
	Killed
)
