package main

import (
	"reflect"
	"testing"
)

func TestCountNums(t *testing.T) {
	inputs := sequence([]int{1, 2, 3, 4, 5, 5})
	expected := map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 2, 6: 0}
	actual := CountNums(inputs)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
