package mindiff

import "testing"

func TestMinDiffOfArrays(t *testing.T) {
	A := []float64{1, 5, 16, 20, 57, 88}
	B := []float64{10, 12, 50, 70}

	ans := MinDiffOfArrays(A, B)
	if ans != 4 {
		t.Errorf("ans == %f; want 4", ans)
	}
}
