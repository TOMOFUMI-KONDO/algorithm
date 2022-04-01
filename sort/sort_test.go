package sort

import "testing"

var sorts = []Sort{BubbleSort, QuickSort}
var tests = []struct {
	in   []int
	want []int
}{
	{in: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
	{in: []int{5, 3, 9, 7, 6, 1}, want: []int{1, 3, 5, 6, 7, 9}},
	{in: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
}

func TestBubbleSort(t *testing.T) {

	for _, sort := range sorts {
		for _, test := range tests {
			sort(&test.in)

			if !same(test.in, test.want) {
				t.Errorf("l = %v; want [1, 3, 5, 6, 7, 9]", test.in)
			}
		}

	}
}

func same(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
