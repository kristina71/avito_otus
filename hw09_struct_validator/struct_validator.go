package hw09structvalidator

import (
	"reflect"
	"sync"
)

type validationResult struct {
	errs ValidationErrors
}

var errsPool = &sync.Pool{
	New: func() interface{} {
		return &validationResult{
			errs: make(ValidationErrors, 0, 16),
		}
	},
}

func validateStruct(sr structRules, value reflect.Value) error {
	result := errsPool.Get().(*validationResult)
	defer func() {
		result.errs = result.errs[:0]
		errsPool.Put(result)
	}()

	for i := 0; i < value.NumField(); i++ {
		rules := sr[i]
		if rules == nil {
			continue
		}
		result.errs = rules.validate(result.errs, value.Field(i))
	}

	if len(result.errs) == 0 {
		return nil
	}
	errs := make(ValidationErrors, len(result.errs))
	copy(errs, result.errs)
	return errs
}
