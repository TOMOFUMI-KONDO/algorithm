package sort

import (
	"sync"
)

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
		if il == mid {
			sorted[i] = (*S)[ir]
			ir++
		} else if ir == r {
			sorted[i] = (*S)[il]
			il++
		} else if (*S)[il] > (*S)[ir] {
			sorted[i] = (*S)[ir]
			ir++
		} else {
			sorted[i] = (*S)[il]
			il++
		}
	}

	for i, s := range sorted {
		(*S)[l+i] = s
	}
}

func MergeSortConcurrent(S *[]int) {
	mergeSortConcurrent(S, 0, len(*S))
}

func mergeSortConcurrent(S *[]int, l, r int) {
	if r-l == 1 {
		return
	}

	mid := (l + r) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mergeSort(S, l, mid)
	}()
	go func() {
		defer wg.Done()
		mergeSort(S, mid, r)
	}()

	wg.Wait()

	sorted := make([]int, r-l)
	il, ir := l, mid
	for i := 0; i < r-l; i++ {
		if il == mid {
			sorted[i] = (*S)[ir]
			ir++
		} else if ir == r {
			sorted[i] = (*S)[il]
			il++
		} else if (*S)[il] > (*S)[ir] {
			sorted[i] = (*S)[ir]
			ir++
		} else {
			sorted[i] = (*S)[il]
			il++
		}
	}

	for i, s := range sorted {
		(*S)[l+i] = s
	}
}
