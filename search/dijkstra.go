package search

import "github.com/amclees/go-practice/graph"

func Dijkstra(g graph.Graph, start, end int) []int {
	q := dqueue{s: make([]pair, 0)}
	_ = q.set(start, 0)

	prev := make(map[int]int)
	done := make(map[int]bool)

	for p, ok := q.next(); ok; p, ok = q.next() {
		if done[p.n] {
			continue
		}
		done[p.n] = true
		for _, edge := range g.Edges(p.n) {
			if q.set(edge[0], p.w+edge[1]) {
				prev[edge[0]] = p.n
			}
		}
	}

	return path(prev, start, end)
}
