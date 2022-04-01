package bubblesort

import "testing"

func TestBubbleSort(t *testing.T) {
	l := [...]int{5, 3, 9, 7, 6, 1}
	sort(l[:])

	/*
		Array values are comparable if values of the array element type are comparable.
		Two array values are equal if their corresponding elements are equal.

		from: https://go.dev/ref/spec#Comparison_operators
	*/
	if l != [...]int{1, 3, 5, 6, 7, 9} {
		t.Errorf("l = %v; want [1, 3, 5, 6, 7, 9]", l)
	}
}
