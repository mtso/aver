package example

import (
	"testing"
	"github.com/mtso/aver"
)

func TestAdd(t *testing.T) {
	aver := aver.New(t)
	aver(add(2, 3)).ToEqual(5)
}

func TestSub(t *testing.T) {
	aver.New(t)(sub(2, 1)).ToEqual(1)
}

func TestSlice(t *testing.T) {
	aver := aver.New(t)
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	aver(slice1).ToEqual(slice2)
}

func TestParity(t *testing.T) {
	aver := aver.New(t)
	isEven := aver(true)
	actual1 := 1 % 2 == 0 // false
	actual2 := 2 % 2 == 0 // true
	isEven.ToNotEqual(actual1)
	isEven.ToEqual(actual2)
}