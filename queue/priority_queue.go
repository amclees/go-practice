package queue

import (
	"github.com/amclees/go-practice/heap"
	"github.com/amclees/go-practice/hmap"
)

type PriorityQueue struct {
	m hmap.Map
	h heap.MaxHeap
}

type intKey int

func (key intKey) Hash() int {
	return int(key)
}

func (q *PriorityQueue) Init(cap int) {
	q.m.Init(cap)
	q.h.Init(cap)
}

func (q *PriorityQueue) Enqueue(p int, d interface{}) {
	k := hmap.Key(intKey(p))
	w := d
	q.m.Put(&k, &w)
	q.h.Add(p)
}

func (q *PriorityQueue) Dequeue() interface{} {
	k := hmap.Key(intKey(q.h.Extract()))
	_, val := q.m.Get(&k)
	return (*val).(interface{})
}
