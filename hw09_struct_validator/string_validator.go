package hw09structvalidator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type stringRule interface {
	validate(value string) error
}

type stringRules struct {
	field string
	vKind validateKind
	rules []stringRule
}

type strLen struct {
	cond int
}

type strIn struct {
	cond []string
}

func (r *stringRules) fieldName() string {
	return r.field
}

func (r *stringRules) validate(errs ValidationErrors, value reflect.Value) ValidationErrors {
	switch r.vKind {
	case validateRegular:
		return r.validateRegular(errs, value)
	case validateSlice:
		return r.validateSlice(errs, value)
	default:
		return r.validateSlice(errs, value)
	}
}

func (r *stringRules) validateRegular(errs ValidationErrors, value reflect.Value) ValidationErrors {
	val := value.String()
	for _, rule := range r.rules {
		err := rule.validate(val)
		if err != nil {
			errs = append(errs, ValidationError{Field: r.field, Err: err})
		}
	}
	return errs
}

func (r *stringRules) validateSlice(errs ValidationErrors, value reflect.Value) ValidationErrors {
	for i := 0; i < value.Len(); i++ {
		errs = r.validateRegular(errs, value.Index(i))
	}
	return errs
}

func newStrLen(value string) (*strLen, error) {
	val, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return nil, err
	}
	return &strLen{cond: int(val)}, nil
}

func (s strLen) validate(value string) error {
	if len(value) == s.cond {
		return nil
	}
	return &ErrStrLen{len(value), s.cond}
}

func newStrRegexp(value string) (*strRegexp, error) {
	rg, err := regexp.Compile(value)
	if err != nil {
		return nil, err
	}
	return &strRegexp{cond: rg}, nil
}

func newStrIn(value string) *strIn {
	val := strings.Split(value, ",")
	return &strIn{cond: val}
}

func (s strRegexp) validate(value string) error {
	if s.cond.MatchString(value) {
		return nil
	}
	return &ErrStrRegexp{value, s.cond}
}

func (s strIn) validate(value string) error {
	if stringContains(s.cond, value) {
		return nil
	}
	return &ErrStrIn{value, s.cond}
}
