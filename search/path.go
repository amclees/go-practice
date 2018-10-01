package search

func path(prev map[int]int, start int, end int) []int {
	p := make([]int, 0)
	ok := true
	for current := end; current != start; current, ok = prev[current] {
		if !ok {
			return []int{}
		}
		p = append(p, current)
	}
	p = append(p, start)

	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}

	return p
}
