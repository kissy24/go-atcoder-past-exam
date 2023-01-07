package main

import (
	"reflect"
	"testing"
)

func TestGenerateDiffs(t *testing.T) {
	sales := []int{9, 10, 100, 10, 1}
	expected := []int{1, 90, -90, -9}
	actual := GenerateDiffs(sales)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestFormat(t *testing.T) {
	params := []struct {
		diff     int
		expected string
	}{
		{diff: 10, expected: "up 10"},
		{diff: -199, expected: "down 199"},
		{diff: 0, expected: "stay"},
	}
	for _, p := range params {
		actual := Format(p.diff)
		if p.expected != actual {
			t.Errorf("Expected: %v, Actual: %v", p.expected, actual)
		}
	}
}
