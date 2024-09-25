package errpkg

import "errors"

var (
	ErrChecklistNotFound = errors.New("checklist not found")
)
