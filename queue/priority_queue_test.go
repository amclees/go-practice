package queue

import "testing"

func TestPriorityQueue(t *testing.T) {
	q := PriorityQueue{}
	q.Init(10)

	q.Enqueue(5, 10)
	q.Enqueue(6, 5)
	q.Enqueue(3, 4)

	expected := []int{5, 10, 4}
	for i := range expected {
		d := q.Dequeue()
		if d != expected[i] {
			t.Errorf("Expected dequeue to return %d, was %d", expected[i], d)
		}
	}
}
