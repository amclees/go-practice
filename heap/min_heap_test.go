package heap

import "testing"

func TestMinHeap(t *testing.T) {
	m := MinHeap{}

	m.Add(5)
	m.Add(3)
	m.Add(10)
	m.Add(7)
	m.Add(6)

	s := make([]int, 5)
	for i := range s {
		s[i] = m.Extract()
	}

	b := []int{3, 5, 6, 7, 10}
	for i := range s {
		if s[i] != b[i] {
			t.Errorf("Expected s[i] to be %d, got %d", b[i], s[i])
		}
	}
}
