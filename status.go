package main

type Status int

const (
	Todo  Status = 0
	Doing Status = 1
	Done  Status = 2
)

func (status Status) String() string {
	names := [...]string{
		"To do",
		"Doing",
		"Done"}

	if status < Todo || status > Done {
		return "Unknown"
	}

	return names[status]
}
