package errs

import "errors"

var (
	ErrTodoNotFound = errors.New("todo not found")
)
