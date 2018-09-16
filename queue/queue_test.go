package queue

import "testing"

func TestQueue(t *testing.T) {
	q := Queue{}

	_, success := q.Dequeue()
	if success {
		t.Errorf("Expected Dequeue to fail on empty queue")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	expected := []int64{1, 2, 3}
	for i := range expected {
		found, _ := q.Dequeue()
		if found != expected[i] {
			t.Errorf("Expected dequeue to return %d, got %d", expected[i], found)
		}
	}
}
