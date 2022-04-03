package search

// SimpleSearch returns index of n in the list.
// If n is not in the list, returns -1.
func SimpleSearch(list []int, n int) int {
	for i, l := range list {
		if l == n {
			return i
		}
	}

	return -1
}
