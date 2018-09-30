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

	p := make([]int, 0)
	ok := true
	for current := end; current != start; current, ok = prev[current] {
		if !ok {
			return []int{}
		}
		p = append(p, current)
	}
	p = append(p, start)

	for i, j := 0, len(p) - 1; i < j; i, j = i + 1, j - 1 {
		p[i], p[j] = p[j], p[i]
	}

	return p
}
