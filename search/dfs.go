package search

import "github.com/amclees/go-practice/graph"

func DFS(g graph.Graph, start, end int) []int {
	st := []int{start}

	prev := make(map[int]int, g.NodeCount())
	for n := st[0]; len(st) != 0; n, st = st[len(st)-1], st[:len(st)-1] {
		for _, edge := range g.Edges(n) {
			if _, ok := prev[edge[0]]; ok {
				continue
			}
			prev[edge[0]] = n
			st = append(st, edge[0])
		}
	}

	return path(prev, start, end)
}
