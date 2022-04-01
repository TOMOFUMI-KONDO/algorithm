package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	sorts := []Sort{BubbleSort, QuickSort}

	for _, sort := range sorts {
		l := []int{5, 3, 9, 7, 6, 1}
		sort(&l)

		if !same(l, []int{1, 3, 5, 6, 7, 9}) {
			t.Errorf("l = %v; want [1, 3, 5, 6, 7, 9]", l)
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
