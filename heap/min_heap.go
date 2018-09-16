package heap

type MinHeap struct {
	s []int64
}

func (m *MinHeap) Init(cap int) {
	m.s = make([]int64, 0, cap)
}

func (m *MinHeap) Add(d int64) {
	i := len(m.s)
	m.s = append(m.s, d)
	s := m.s
	for {
		ni := i / 2
		if s[ni] > s[i] {
			a := s[ni]
			s[ni] = s[i]
			s[i] = a
		}

		if i == 0 {
			break
		}
		i = ni
	}
}

func (m *MinHeap) Extract() int64 {
	s := m.s
	d := s[0]
	s[0] = s[len(s)-1]
	m.s = s[:len(s)-1]

	m.heapify(0)

	return d
}

func (m *MinHeap) heapify(i int) {
	s := m.s
	l := (i * 2) + 1
	r := l + 1
	len := len(s)

	swap := -1
	if l < len && s[l] < s[i] {
		swap = l
	}
	if r < len && s[r] < s[i] && (swap == -1 || s[swap] > s[r]) {
		swap = r
	}

	if swap != -1 {
		a := s[swap]
		s[swap] = s[i]
		s[i] = a

		m.heapify(swap)
	}
}
