package main

import (
	"reflect"
	"testing"
)

func TestToSeq(t *testing.T) {
	inputs := []string{"3", "6", "19", "20", "7"}
	expected := sequence([]int{3, 6, 19, 20, 7})
	actual, err := ToSeq(inputs)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

}
