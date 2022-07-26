package hw09structvalidator

import (
	"reflect"
	"sync/atomic"
)

type structBuffer struct {
	values atomic.Value
}

func newStructBuffer() *structBuffer {
	values := make(map[reflect.Type]structRules)
	cache := &structBuffer{}
	cache.values.Store(values)
	return cache
}

func (sc *structBuffer) lookup(value reflect.Type) (structRules, bool) {
	v, ok := sc.values.Load().(map[reflect.Type]structRules)[value]
	return v, ok
}

func (sc *structBuffer) add(value reflect.Type, rules structRules) {
	values := sc.values.Load().(map[reflect.Type]structRules)
	newValues := make(map[reflect.Type]structRules, len(values)+1)
	for k, v := range values {
		newValues[k] = v
	}
	newValues[value] = rules
	sc.values.Store(newValues)
}
