package main

import "testing"

func TestSortByDoubleCamel(t *testing.T) {
	params := []struct {
		inp      string
		expected string
	}{
		{inp: "FisHDoGCaTAAAaAAbCAC", expected: "AAAaAAbCACCaTDoGFisH"},
		{inp: "AAAAAjhfgaBCsahdfakGZsZGdEAA", expected: "AAAAAAAjhfgaBCsahdfakGGdEZsZ"},
	}

	for _, p := range params {
		actual := SortByDoubleCamel(p.inp)
		if p.expected != actual {
			t.Errorf("Expected: %v, Actual: %v", p.expected, actual)
		}
	}
}
