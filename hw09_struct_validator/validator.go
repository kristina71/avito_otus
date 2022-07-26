package hw09structvalidator

import (
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var b strings.Builder
	b.WriteString("validation errors: ")
	for _, err := range v {
		b.WriteString(err.Field)
		b.WriteString(": ")
		b.WriteString(err.Err.Error())
		b.WriteString("; ")
	}
	return b.String()
}

var cache = newStructBuffer()

func Validate(v interface{}) error {
	if v == nil {
		return nil
	}

	value := reflect.Indirect(reflect.ValueOf(v))
	if value.Kind() != reflect.Struct {
		return &ErrIncorrectUse{reason: IncorrectKind, kind: value.Kind()}
	}

	structType := value.Type()
	rules, ok := cache.lookup(structType)
	if ok {
		return validateStruct(rules, value)
	}

	rules, err := parseStructRules(value)
	if err != nil {
		return err
	}
	cache.add(structType, rules)
	return validateStruct(rules, value)
}
