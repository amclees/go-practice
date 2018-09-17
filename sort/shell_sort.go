package sort

var gaps [9]int = [9]int{1750, 701, 301, 132, 57, 23, 10, 4, 1}

func Shellsort(s []int64) {
	l := len(s)
	i := 0
	switch {
	case l > gaps[0]:
		i = 0
	case l > gaps[1]:
		i = 1
	case l > gaps[2]:
		i = 2
	case l > gaps[3]:
		i = 3
	case l > gaps[4]:
		i = 4
	case l > gaps[5]:
		i = 5
	case l > gaps[6]:
		i = 6
	case l > gaps[7]:
		i = 7
	case l > gaps[8]:
		i = 8
	}

	for ; i < len(gaps); i++ {
		for offset := 0; offset < gaps[i]; offset++ {
			for j := offset; j < l; j += gaps[i] {
				for k := j - gaps[i]; k >= 0; k -= gaps[i] {
					n := k + gaps[i]
					if s[n] < s[k] {
						a := s[n]
						s[n] = s[k]
						s[k] = a
					} else {
						break
					}
				}
			}
		}
	}
}
