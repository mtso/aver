# aver [![GoDoc](https://godoc.org/github.com/mtso/aver?status.svg)](https://godoc.org/github.com/mtso/aver)

Assertion library for Go.

## Usage

Package aver adds assertion functions for Go. Its usage is inspired by the "expect" assertion pattern.

Usage
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
		aver(actual).ToEqual(expected)
	}

## Aver is Expect

Expect assertions allow you to compare two 