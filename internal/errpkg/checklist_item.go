package errpkg

import "errors"

var (
	ErrChecklistItemNotFound = errors.New("checklist item not found")
)
