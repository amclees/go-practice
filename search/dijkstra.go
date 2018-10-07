package search

import "github.com/amclees/go-practice/graph"

func Dijkstra(g graph.Graph, start, end int) []int {
	q := dqueue{s: make([]pair, 0)}
	_ = q.set(start, 0)

	prev := make(map[int]int)
	done := make(map[int]bool)

	for p, ok := q.next(); ok; p, ok = q.next() {
		for _, edge := range g.Edges(p.n) {
			if _, ok := done[edge[0]]; !ok && q.set(edge[0], p.w+edge[1]) {
				done[p.n] = true
				prev[edge[0]] = p.n
			}
		}
	}

	return path(prev, start, end)
}
