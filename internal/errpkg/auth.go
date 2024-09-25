package errpkg

import "errors"

var (
	ErrInvalidEmailOrPassword = errors.New("email or password invalid")
)
