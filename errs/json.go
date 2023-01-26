package errs

import "encoding/json"

var (
	ErrJsonSyntax    *json.SyntaxError
	ErrUnmarshalType *json.UnmarshalTypeError
)
