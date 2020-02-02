package api

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrResourceNotFound = Error("resource not found")
)