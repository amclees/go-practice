package queue

type Queue struct {
	next *Node
	end  *Node
	size int
}

type Node struct {
	val  int64
	next *Node
}

func (q *Queue) Enqueue(d int64) {
	q.size += 1

	n := Node{val: d}

	if q.size == 1 {
		q.next = &n
	} else {
		q.end.next = &n
	}

	q.end = &n
}

func (q *Queue) Dequeue() (int64, bool) {
	if q.size == 0 {
		return 0, false
	} else if q.size == 1 {
		q.end = nil
	}
	q.size -= 1
	d := q.next.val
	q.next = q.next.next
	return d, true
}
