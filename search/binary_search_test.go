package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		list []int
		n    int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 20, 34, 49, 510, 6000}, 1, 0},
		{[]int{1, 2, 5, 7, 8}, 3, -1},
	}

	for _, test := range tests {
		res := BinarySearch(test.list, test.n)
		if res != test.want {
			t.Errorf("res == %d; want %d", res, test.want)
		}
	}
}
