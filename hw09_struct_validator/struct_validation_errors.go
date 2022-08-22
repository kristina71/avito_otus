package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
)

type IncorrectUseCase int

const (
	IncorrectKind      IncorrectUseCase = iota // incorrect type (not a struct)
	IncorrectFieldType                         // unsupported type (not int, string, []int, []string)
	UnknownRule                                // unknown validation rule
	IncorrectCondition                         // can not parse struct
)

type ErrIncorrectUse struct {
	reason IncorrectUseCase
	kind   reflect.Kind
	field  string
	rule   string
	err    error
}

func (e *ErrIncorrectUse) Error() string {
	switch e.reason {
	case IncorrectKind:
		return fmt.Sprintf("function only accepts structs; got %s", e.kind)
	case IncorrectFieldType:
		return fmt.Sprintf("field `%s` has unsupported type %s", e.field, e.kind)
	case UnknownRule:
		return fmt.Sprintf("field `%s` has unknown rule `%s`", e.field, e.rule)
	case IncorrectCondition:
		return fmt.Sprintf("field `%s` has incorrect condition: `%s` for rule `%s`", e.field, e.err, e.rule)
	default:
		return ""
	}
}

func (e *ErrIncorrectUse) Is(target error) bool {
	var err *ErrIncorrectUse
	return errors.As(target, &err)
}

func (e *ErrIncorrectUse) Unwrap() error {
	return e.err
}
