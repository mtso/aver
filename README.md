# aver [![GoDoc](https://godoc.org/github.com/mtso/aver?status.svg)](https://godoc.org/github.com/mtso/aver)

Assertion library for Go.

## Usage

Example usage in a *_test.go file:
```go
import (
	"testing"
	"github.com/mtso/aver"
)

func TestExample(t *testing.T) {
	// Create a new Averrer instance
	aver := aver.New(t)

	// FAIL the test with Errorf
	expected := "expectedString"
	actual := "actualString"
	aver(expected).ToEqual(actual)
	
	// Store the returned Averrable in a variable
	isEven := aver(true)
	actual1 := 1 % 2 == 0 // false
	actual2 := 2 % 2 == 0 // true
	isEven.ToNotEqual(actual1)
	isEven.ToEqual(actual2)
}
```

## Aver is Expect

Aver (Expect) assertions allow you to compare two values in a test.
aver.New(t *testing.T) returns a new Averrer function.
The Averrer function accepts an "expected" value as a parameter and returns
an Aver object that conforms to the Averrable interface. The Aver
object stores the expected value and a reference to the testing.T instance.
The Aver object can then be called with an Averrable function to compare
the stored "expected" value with an "actual" value passed into the 
Averrable function.