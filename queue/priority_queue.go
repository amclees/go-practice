package queue

import "github.com/amclees/go-practice/heap"

type PriorityQueue struct {
	h heap.MaxHeap
}

type pair struct {
	key int
	val interface{}
}

func (p pair) Key() int {
	return p.key
}

func (p pair) Val() interface{} {
	return p.val
}

func (q *PriorityQueue) Init(cap int) {
	q.h.Init(cap)
}

func (q *PriorityQueue) Enqueue(p int, d interface{}) {
	q.h.Add(pair{p, d})
}

func (q *PriorityQueue) Dequeue() interface{} {
	return q.h.Extract().Val()
}
