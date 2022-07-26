package hw09structvalidator

import "reflect"

const validatorTag = "validate"

type structRules []fieldRules

type fieldRules interface {
	fieldName() string
	validate(errs ValidationErrors, value reflect.Value) ValidationErrors
}

type validateKind int

const (
	validateRegular validateKind = iota
	validateSlice
)
