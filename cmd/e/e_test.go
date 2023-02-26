package main

import (
	"reflect"
	"testing"
)

func TestToInts(t *testing.T) {
	inputs := []string{"1", "2", "3"}
	expected := []int{1, 2, 3}
	actual := ToInts(inputs)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
