package sort

func QuickSort(lp *[]int) {
	l := *lp

	if len(l) == 1 {
		return
	}

	th := threshold(l)
	var left, right []int
	for _, x := range l {
		if x <= th {
			left = append(left, x)
		} else {
			right = append(right, x)
		}
	}

	QuickSort(&left)
	QuickSort(&right)

	merged := append(left, right...)
	*lp = merged
}

func threshold(l []int) int {
	return (l[0] + l[1]) / 2
}
