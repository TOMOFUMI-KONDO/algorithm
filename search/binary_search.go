package search

// BinarySearch returns index of n in the list.
// If n is not in the list, returns -1.
func BinarySearch(list []int, n int) int {
	ok := 0
	ng := len(list)

	for ng-ok > 1 {
		mid := (ok + ng) / 2
		if list[mid] <= n {
			ok = mid
		} else {
			ng = mid
		}
	}

	if list[ok] == n {
		return ok
	} else {
		return -1
	}
}
