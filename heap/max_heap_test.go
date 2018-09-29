package heap

import "testing"

func TestMaxHeap(t *testing.T) {
	m := MaxHeap{}

	m.Add(5)
	m.Add(3)
	m.Add(10)
	m.Add(7)
	m.Add(6)

	s := make([]int, 5)
	for i := range s {
		s[i] = m.Extract()
	}

	b := []int{10, 7, 6, 5, 3}
	for i := range s {
		if s[i] != b[i] {
			t.Errorf("Expected s[i] to be %d, got %d", b[i], s[i])
		}
	}
}
