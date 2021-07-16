package subset_sum_slow

import (
	"reflect"
	"testing"
)

func TestSubsetSumSlow_normal(t *testing.T) {
	yn, x := SubsetSumSlow([]int{3, 7, 5, 8, 2}, 11)

	if yn != Y || reflect.DeepEqual(x, []int{1, 0, 0, 1, 0}) {
		t.Errorf("")
	}
}
