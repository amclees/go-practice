package heap

import "testing"

type intVal int

func (i intVal) Key() int {
	return int(i)
}

func (i intVal) Val() interface{} {
	return int(i)
}

func TestMaxHeap(t *testing.T) {
	m := MaxHeap{}

	m.Add(intVal(5))
	m.Add(intVal(3))
	m.Add(intVal(10))
	m.Add(intVal(7))
	m.Add(intVal(6))

	s := make([]int, 5)
	for i := range s {
		s[i] = m.Extract().Key()
	}

	b := []int{10, 7, 6, 5, 3}
	for i := range s {
		if s[i] != b[i] {
			t.Errorf("Expected s[i] to be %d, got %d", b[i], s[i])
		}
	}
}
