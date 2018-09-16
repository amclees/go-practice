package heap

import "testing"

func TestMaxHeap(t *testing.T) {
	m := MaxHeap{}

	m.Add(int64(5))
	m.Add(int64(3))
	m.Add(int64(10))
	m.Add(int64(7))
	m.Add(int64(6))

	s := make([]int64, 5)
	for i := range s {
		s[i] = m.Extract()
	}

	b := []int64{10, 7, 6, 5, 3}
	for i := range s {
		if s[i] != b[i] {
			t.Errorf("Expected s[i] to be %d, got %d", b[i], s[i])
		}
	}
}
