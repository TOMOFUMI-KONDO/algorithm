package search

// BitFullSearch returns all subsets of list having n elements.
func BitFullSearch(n int) [][]int {
	result := make([][]int, 1<<n)

	for mask := 0; mask < 1<<n; mask++ {
		var l []int

		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				l = append(l, i)
			}
		}

		result[mask] = l
	}

	return result
}
