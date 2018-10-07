package heap

type MaxHeap struct {
	s []Interface
}

func (m *MaxHeap) Init(cap int) {
	m.s = make([]Interface, 0, cap)
}

func (m *MaxHeap) Add(d Interface) {
	i := len(m.s)
	m.s = append(m.s, d)
	s := m.s
	for {
		ni := i / 2
		if s[ni].Key() < s[i].Key() {
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

func (m *MaxHeap) Extract() Interface {
	s := m.s
	d := s[0]
	s[0] = s[len(s)-1]
	m.s = s[:len(s)-1]

	m.Heapify(0)

	return d
}

func (m *MaxHeap) Heapify(i int) {
	s := m.s
	l := (i * 2) + 1
	r := l + 1
	len := len(s)

	swap := -1
	if l < len && s[l].Key() > s[i].Key() {
		swap = l
	}
	if r < len && s[r].Key() > s[i].Key() && (swap == -1 || s[swap].Key() < s[r].Key()) {
		swap = r
	}

	if swap != -1 {
		a := s[swap]
		s[swap] = s[i]
		s[i] = a

		m.Heapify(swap)
	}
}
