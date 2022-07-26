package hw09structvalidator

import "fmt"

type ErrIntIn struct {
	Value int64
	Cond  []int64
}

// min

func (e *ErrIntIn) Error() string {
	return fmt.Sprintf("number %d is not included in the specified set %v", e.Value, e.Cond)
}

type ErrIntMin struct {
	Value int64
	Cond  int64
}

func (e *ErrIntMin) Error() string {
	return fmt.Sprintf("number %d is less than specified %d", e.Value, e.Cond)
}

// max

type ErrIntMax struct {
	Value int64
	Cond  int64
}

func (e *ErrIntMax) Error() string {
	return fmt.Sprintf("number %d is greater than specified %d", e.Value, e.Cond)
}
