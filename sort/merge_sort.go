package sort

func MergeSort(S *[]int) {
	mergeSort(S, 0, len(*S))
}

func mergeSort(S *[]int, l, r int) {
	if r-l == 1 {
		return
	}

	mid := (l + r) / 2

	mergeSort(S, l, mid)
	mergeSort(S, mid, r)

	sorted := make([]int, r-l)
	il, ir := l, mid
	for i := 0; i < r-l; i++ {
		if il == mid || (*S)[il] > (*S)[ir] {
			sorted[i] = (*S)[ir]
			ir++
		} else {
			sorted[i] = (*S)[il]
			il++
		}
	}
}
