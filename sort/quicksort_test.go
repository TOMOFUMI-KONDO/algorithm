package sort

import (
	"testing"

	alg "github.com/TOMOFUMI-KONDO/algorithm"
)

func TestQuickSort(t *testing.T) {
	for _, test := range tests {
		QuickSort(&test.in)

		if !alg.SameIntSlice(test.in, test.want) {
			t.Errorf("result = %v; want %v", test.in, test.want)
		}
	}
}
