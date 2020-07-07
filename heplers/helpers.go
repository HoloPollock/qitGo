package helpers

//Status Repersentes the status of a item
type Status int

const (
	Added Status = iota
	Changed
	Deleted
)
