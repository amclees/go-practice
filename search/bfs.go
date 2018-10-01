package search

import (
	"github.com/amclees/go-practice/graph"
	"github.com/amclees/go-practice/queue"
)

func BFS(g graph.Graph, start, end int) []int {
	q := queue.Queue{}
	q.Enqueue(start)

	prev := make(map[int]int)
	for n, ok := q.Dequeue(); ok; n, ok = q.Dequeue() {
		for _, edge := range g.Edges(n) {
			_, ok = prev[edge[0]]
			if ok {
				continue
			}
			prev[edge[0]] = n
			q.Enqueue(edge[0])
		}
	}

	return path(prev, start, end)
}
