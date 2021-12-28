package main

import (
	"math"
)

func MinDiffOfArrays(A []float64, B []float64) float64 {
	i := 0
	j := 0

	var ans = math.Inf(1)
	for i < len(A) && j < len(B) {
		a := A[i]
		b := B[j]

		if a == b {
			ans = 0
			break
		}

		ans = math.Min(math.Abs(float64(a-b)), ans)

		if a < b {
			i++
		} else {
			j++
		}
	}

	return ans
}
