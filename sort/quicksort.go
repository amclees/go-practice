package sort

func Quicksort(s []int64) {
	if len(s) < 2 {
		return
	}

	p := partition(s)
	Quicksort(s[:p+1])
	Quicksort(s[p+1:])
}

func partition(s []int64) int {
	pivot := findPivot(s)
	i := 0
	j := len(s) - 1
	for {
		for ; s[i] < pivot; i++ {
		}

		for ; s[j] > pivot; j-- {
		}

		if i >= j {
			return j
		}

		t := s[i]
		s[i] = s[j]
		s[j] = t

		i += 1
		j -= 1
	}
}

func findPivot(s []int64) int64 {
	return s[0]
}
