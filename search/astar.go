package search

import "github.com/amclees/go-practice/graph"

func AStar(h func(graph.Graph, int, int) int) func(graph.Graph, int, int) []int {
	return func(g graph.Graph, start, end int) []int {
		q := dqueue{s: make([]pair, 0)}
		_ = q.set(start, 0)

		prev := make(map[int]int)
		done := make(map[int]bool)

		known_score := make(map[int]int)

		for p, ok := q.next(); ok; p, ok = q.next() {
			done[p.n] = true
			here := known_score[p.n]
			for _, edge := range g.Edges(p.n) {
				if _, ok := done[edge[0]]; ok {
					continue
				}
				thru_score := here + edge[1]

				if val, ok := known_score[thru_score]; ok && val < thru_score {
					continue
				}

				est := thru_score + h(g, edge[0], end)

				if !q.set(edge[0], est) {
					continue
				}
				prev[edge[0]] = p.n
				known_score[edge[0]] = thru_score
			}
		}

		return path(prev, start, end)
	}
}
