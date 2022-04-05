package search

import (
	"testing"

	"github.com/TOMOFUMI-KONDO/algorithm"
)

var tests = []struct {
	n    int
	want [][]int
}{
	{3, [][]int{{}, {0}, {1}, {0, 1}, {2}, {0, 2}, {1, 2}, {0, 1, 2}}},
	{4, [][]int{{}, {0}, {1}, {0, 1}, {2}, {0, 2}, {1, 2}, {0, 1, 2}, {3}, {0, 3}, {1, 3}, {0, 1, 3}, {2, 3}, {0, 2, 3}, {1, 2, 3}, {0, 1, 2, 3}}},
}

func TestBitFullSearch(t *testing.T) {
	for _, test := range tests {
		res := BitFullSearch(test.n)
		if !algorithm.SameIntSliceSlice(res, test.want) {
			t.Errorf("res == %v; want %v", res, test.want)
		}
	}
}
