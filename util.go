package algorithm

func SameIntSlice(a, b []int) bool {
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

func SameIntSliceSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !SameIntSlice(a[i], b[i]) {
			return false
		}
	}

	return true
}
