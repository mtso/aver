package aver

import (
	"reflect"
	"testing"
)

type Averrable interface {
	ToEqual(interface{})
	ToNotEqual(interface{})
}

func New(t *testing.T) func(actual interface{}) Averrable {
	return func(actual interface{}) Averrable {
		return Averrer{
			actual,
			t,
		}
	}
}

type Averrer struct {
	actual interface{}
	*testing.T
}

func (a Averrer) ToEqual(expected interface{}) {
	eq := reflect.DeepEqual(a.actual, expected)
	if !eq {
		a.Errorf("Expected %q to equal %q", a.actual, expected)
	}
}

func (a Averrer) ToNotEqual(expected interface{}) {
	if reflect.DeepEqual(a.actual, expected) {
		a.Errorf("Expected %q to NOT equal %q", a.actual, expected)
	}
}
