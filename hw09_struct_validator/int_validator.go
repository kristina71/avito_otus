package hw09structvalidator

import (
	"reflect"
	"strconv"
	"strings"
)

type intIn struct {
	cond []int64
}

func newIntIn(value string) (*intIn, error) {
	val, err := stringsToInts64(strings.Split(value, ","))
	if err != nil {
		return nil, err
	}
	return &intIn{cond: val}, nil
}

type intMax struct {
	cond int64
}

func newIntMax(value string) (*intMax, error) {
	val, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return nil, err
	}
	return &intMax{cond: val}, nil
}

type intMin struct {
	cond int64
}

func newIntMin(value string) (*intMin, error) {
	val, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return nil, err
	}
	return &intMin{cond: val}, nil
}

func (s intMin) validate(value int64) error {
	if value >= s.cond {
		return nil
	}
	return &ErrIntMin{value, s.cond}
}

func (s intMax) validate(value int64) error {
	if value <= s.cond {
		return nil
	}
	return &ErrIntMax{value, s.cond}
}

func (s intIn) validate(value int64) error {
	if intContains(s.cond, value) {
		return nil
	}
	return &ErrIntIn{value, s.cond}
}

type intRule interface {
	validate(value int64) error
}

type intRules struct {
	field string
	vKind validateKind
	rules []intRule
}

func (r *intRules) fieldName() string {
	return r.field
}

func (r *intRules) validate(errs ValidationErrors, value reflect.Value) ValidationErrors {
	switch r.vKind {
	case validateRegular:
		return r.validateRegular(errs, value)
	case validateSlice:
		return r.validateSlice(errs, value)
	default:
		return r.validateSlice(errs, value)
	}
}

func (r *intRules) validateRegular(errs ValidationErrors, value reflect.Value) ValidationErrors {
	val := value.Int()
	for _, rule := range r.rules {
		err := rule.validate(val)
		if err != nil {
			errs = append(errs, ValidationError{Field: r.field, Err: err})
		}
	}
	return errs
}

func (r *intRules) validateSlice(errs ValidationErrors, value reflect.Value) ValidationErrors {
	for i := 0; i < value.Len(); i++ {
		errs = r.validateRegular(errs, value.Index(i))
	}
	return errs
}
