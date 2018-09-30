package search

import (
	"testing"

	"github.com/amclees/go-practice/graph"
)

type testCase struct {
	g                  graph.Graph
	start, end, minLen int
}

func TestBFS(t *testing.T) {
	// BFS does not consider weights and will not necessarily produce minimum paths
	testSearch(t, BFS, false)
}

func testSearch(t *testing.T, search func(graph.Graph, int, int) []int, testMin bool) {
	for i, c := range createCases() {
		p := search(c.g, c.start, c.end)
		if !verifyPath(c.g, p, c.start, c.end) {
			t.Errorf("Case %d: Expected valid path, got %v", i, p)
		}
		if testMin && len(p) != c.minLen {
			t.Errorf("Case %d: Expected path length %d, got %d", i, c.minLen, len(p))
		}
	}
}

func verifyPath(g graph.Graph, p []int, start, end int) bool {
	if len(p) == 0 || p[0] != start || p[len(p)-1] != end {
		return false
	}
	for i := 0; i < len(p)-1; i++ {
		if ok, _ := g.Edge(p[i], p[i+1]); !ok {
			return false
		}
	}
	return true
}

func createCases() []testCase {
	g1 := createGraph()
	g1.AddNode(1, 0)
	g1.AddNode(2, 0)
	g1.AddNode(3, 0)
	g1.AddNode(4, 0)
	g1.AddNode(5, 0)
	g1.AddEdge(1, 4, 20)
	g1.AddEdge(1, 3, 4)
	g1.AddEdge(1, 5, 1)
	g1.AddEdge(3, 2, 2)
	g1.AddEdge(3, 4, 1)
	g1.AddEdge(4, 1, 2)
	g1.AddEdge(5, 1, 1)
	g1.AddEdge(5, 4, 10)

	s := []testCase{{g: g1, start: 1, end: 4, minLen: 3}}

	return s
}

func createGraph() graph.Graph {
	g := graph.PointerGraph{}
	return graph.Graph(&g)
}
