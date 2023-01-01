package main

import (
	"errors"
	"strconv"
	"testing"
)

func TestToDoubledInt(t *testing.T) {
	params := []struct {
		inp      string
		expected int
		err      error
	}{
		{inp: "678", expected: 1356, err: nil},
		{inp: "abc", expected: 0, err: strconv.ErrSyntax},
		{inp: "0x8", expected: 0, err: strconv.ErrSyntax},
		{inp: "012", expected: 24, err: nil},
	}

	for _, p := range params {
		actual, err := ToDoubledInt(p.inp)
		if !errors.Is(err, p.err) {
			t.Errorf("Expected: %v, Actual: %v", p.err, err)
		}
		if p.expected != actual {
			t.Errorf("Expected: %v, Actual: %v", p.expected, actual)
		}
	}
}
