package sort

func BubbleSort(s []int64) {
	for {
		swapped := false
		for i := 0; i < len(s)-1; i++ {
			if s[i+1] < s[i] {
				t := s[i]
				s[i] = s[i+1]
				s[i+1] = t
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}
