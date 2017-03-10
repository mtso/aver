package aver

import "testing"

func TestAver(t *testing.T) {
	aver := New(t)
	aver(4).ToEqual(4)

	averHello := aver("hello")
	averHello.ToEqual("hello")
}
