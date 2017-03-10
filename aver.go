package aver

import (
	"reflect"
	"testing"
)

// Averrable describes functions that can be called on an averred expected value.
type Averrable interface {
	// Compares a passed actual value with the stored expected value for equality.
	ToEqual(actual interface{})
	// Compares a passed actual value with the stored expected value for inequality.
	ToNotEqual(actual interface{})
}

// Averrer produces a new Averrable from an expected value.
// Stores the expected value and returns an Averrable.
type Averrer func(expected interface{}) Averrable

// aver wraps testing.T and stores an expected value for use in Averrable functions.
type aver struct {
	expected interface{}
	*testing.T
}

// Returns a new Averrer function.
func New(t *testing.T) Averrer {
	return func(expected interface{}) Averrable {
		return aver{
			expected,
			t,
		}
	}
}

func (a aver) ToEqual(actual interface{}) {
	eq := reflect.DeepEqual(a.expected, actual)
	if !eq {
		a.Errorf("Expected %q to equal %q", a.expected, actual)
	}
}

func (a aver) ToNotEqual(actual interface{}) {
	if reflect.DeepEqual(a.expected, actual) {
		a.Errorf("Expected %q to NOT equal %q", a.expected, actual)
	}
}
