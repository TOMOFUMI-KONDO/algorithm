package sort

import (
	"testing"

	alg "github.com/TOMOFUMI-KONDO/algorithm"
)

func TestMergeeSort(t *testing.T) {
	for _, test := range tests {
		MergeSort(&test.in)

		if !alg.SameIntSlice(test.in, test.want) {
			t.Errorf("l = %v; want [1, 3, 5, 6, 7, 9]", test.in)
		}
	}

}

func TestMergeSortConcurrent(t *testing.T) {
	for _, test := range tests {
		MergeSortConcurrent(&test.in)

		if !alg.SameIntSlice(test.in, test.want) {
			t.Errorf("l = %v; want [1, 3, 5, 6, 7, 9]", test.in)
		}
	}
}
