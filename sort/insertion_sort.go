package sort

func InsertionSort(s []int64) {
	for i, x := range s {
		for j := i - 1; j >= 0; j-- {
			if x < s[j] {
				y := s[j]
				s[j] = s[j+1]
				s[j+1] = y
			} else {
				break
			}
		}
	}
}
