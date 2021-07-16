package subset_sum_slow

import (
	"reflect"
	"testing"
)

func TestSubsetSumSlowY_normal(t *testing.T) {
	yn, x := SubsetSumSlow([]int{3, 7, 5, 8, 2}, 11)

	if yn != Y || !reflect.DeepEqual(x, []bool{true, false, false, true, false}) {
		t.Errorf("SubsetSumSlow([]int{3, 7, 5, 8, 2}, 11) = %s, %v; want %s, [true, false, false, true, false]", yn, x, Y)
	}
}

func TestSubsetSumSlowN_normal(t *testing.T) {
	yn, x := SubsetSumSlow([]int{3, 7, 5, 8, 2}, 999)

	if yn != N || len(x) != 0 {
		t.Errorf("SubsetSumSlow([]int{3, 7, 5, 8, 2}, 999) = %s, %v; want %s, []", yn, x, N)
	}
}
