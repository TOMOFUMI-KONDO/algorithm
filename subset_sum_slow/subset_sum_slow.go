package subset_sum_slow

type YN string

const (
	Y YN = "Yes"
	N YN = "No"
)

func SubsetSumSlow(a []int, b int) (YN, []bool) {
	lenA := len(a)
	var x []bool
	for i := 0; i < lenA; i++ {
		x = append(x, false)
	}

	for {
		tmp := 0
		for i := 0; i < lenA; i++ {
			if x[i] {
				tmp += a[i]
			}
		}
		if tmp == b {
			return Y, x
		}

		isFull := true
		for i := 0; i < lenA; i++ {
			isFull = x[i]
			if !isFull {
				break
			}
		}
		if isFull {
			return N, nil
		}
		next(x)
	}
}

func next(x []bool) {
	i := len(x) - 1

	for {
		if x[i] {
			x[i] = false
		} else {
			x[i] = true
			break
		}

		i--
	}
}
