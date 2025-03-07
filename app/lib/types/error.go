package types

import "github.com/pkg/errors"

const KeyError = "error"

type Error struct {
	Message string `json:"message"`
}

var _ Type = (*Error)(nil)

func (x *Error) Key() string {
	return KeyError
}

func (x *Error) Sortable() bool {
	return false
}

func (x *Error) Scalar() bool {
	return false
}

func (x *Error) String() string {
	return "error(" + x.Message + ")"
}

func (x *Error) From(v any) any {
	switch t := v.(type) {
	case error:
		return t
	case string:
		return errors.New(t)
	default:
		return invalidInput(x.Key(), t)
	}
}

func (x *Error) Default(string) any {
	return KeyError
}

func NewError(msg string) *Wrapped {
	return Wrap(&Error{Message: msg})
}
