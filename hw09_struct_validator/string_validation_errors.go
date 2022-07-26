package hw09structvalidator

import (
	"fmt"
	"regexp"
)

// len string

type ErrStrLen struct {
	Value int
	Cond  int
}

func (e *ErrStrLen) Error() string {
	return fmt.Sprintf("string length %d is not equal to required %d", e.Value, e.Cond)
}

type ErrStrIn struct {
	Value string
	Cond  []string
}

func (e *ErrStrIn) Error() string {
	return fmt.Sprintf("string `%s` is not included in the specified set %v", e.Value, e.Cond)
}

// regexp string

type strRegexp struct {
	cond *regexp.Regexp
}

type ErrStrRegexp struct {
	Value string
	Cond  *regexp.Regexp
}

func (e *ErrStrRegexp) Error() string {
	return fmt.Sprintf("string `%s` does not match the regexp `%v`", e.Value, e.Cond)
}
