package types

import "errors"

var (
	ErrNotSlice           = errors.New("minecomm: struct field is not a slice")
	ErrMissingLen         = errors.New("minecomm: missing len struct tag where absolutely necessary")
	ErrIncorrectFieldType = errors.New(
		"minecomm: the target field type does not correspond to the one specified in the type tag",
	)
)
