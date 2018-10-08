package search

import (
	"testing"

	"github.com/amclees/go-practice/graph"
)

type testCase struct {
	g                  graph.Graph
	start, end, minLen int
	grid               bool
}

type point struct {
	x, y int
}

func TestBFS(t *testing.T) {
	// BFS does not consider weights and will not necessarily produce minimum paths
	testSearch("BFS", t, BFS, false, false)
}

func TestDFS(t *testing.T) {
	testSearch("DFS", t, DFS, false, false)
}

func TestDijkstra(t *testing.T) {
	testSearch("Dijkstra", t, Dijkstra, true, false)
}

func TestAStar(t *testing.T) {
	testSearch("AStar", t, AStar(h), false, true)
}

func h(g graph.Graph, n1, n2 int) int {
	_, d1 := g.Node(n1)
	_, d2 := g.Node(n2)
	p1, p2 := d1.(point), d2.(point)
	x1, x2 := p1.x, p2.x
	y1, y2 := p1.y, p2.y
	return abs(x1-x2) + abs(y1-y2)
}

func abs(k int) int {
	if k < 0 {
		return k * -1
	}
	return k
}

func testSearch(name string, t *testing.T, search func(graph.Graph, int, int) []int, testMin bool, gridOnly bool) {
	for i, c := range createCases() {
		if gridOnly && !c.grid {
			continue
		}
		p := search(c.g, c.start, c.end)
		if !verifyPath(c.g, p, c.start, c.end) {
			t.Errorf("Case %d (%v): Expected valid path, got %v", i, name, p)
		}

		length := 0
		for j := 0; j < len(p)-1; j++ {
			_, e := c.g.Edge(p[j], p[j+1])
			length += e
		}
		if testMin && length != c.minLen {
			t.Errorf("Case %d (%v): Expected path length %d, got %d (path %v)", i, name, c.minLen, length, p)
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

	g2 := createGraph()
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			g2.AddNode(i*10+j, point{i, j})
		}
	}
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			start := i*10 + j
			points := []point{{i, j + 1}, {i, j - 1}, {i + 1, j}, {i - 1, j}}
			for _, p := range points {
				if p.x > 0 && p.y > 0 && p.x < 5 && p.y < 5 {
					g2.AddEdge(start, p.x*10+p.y, 1)
				}
			}
		}
	}

	s := []testCase{{g: g1, start: 1, end: 4, minLen: 5, grid: false},
		{g: g2, start: 11, end: 44, minLen: 6, grid: true}}

	return s
}

func createGraph() graph.Graph {
	g := graph.PointerGraph{}
	return graph.Graph(&g)
}

func createBenchmarkGraph(n int) graph.Graph {
	nodeMax := translatePoint(point{n, n})
	mg := graph.NewListGraph(nodeMax, 4)
	g := graph.Graph(&mg)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			p := point{i, j}
			g.AddNode(translatePoint(p), p)
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			k := translatePoint(point{i, j})
			points := []point{{i, j + 1}, {i, j - 1}, {i + 1, j}, {i - 1, j}}
			for _, p := range points {
				if p.x > 0 && p.y > 0 && p.x <= n && p.y <= n && (p.x < n/2 || p.y > n/2) {
					g.AddEdge(k, translatePoint(p), 1)
				}
			}
		}
	}

	return g
}

func translateNode(n int) point {
	sx := (n & 0xf0f0f0f0) >> 4
	x := 0
	for ; sx != 0; sx >>= 4 {
		x += sx & 0xf
	}

	sy := n & 0x0f0f0f0f
	y := 0
	for ; sy != 0; sy >>= 4 {
		x += sy & 0xf
	}

	return point{x, y}
}

func translatePoint(p point) int {
	sum := 0
	b1, b2, b3, b4 := p.x&0xf000, p.x&0x0f00, p.x&0x00f0, p.x&0x000f
	sum += b1<<16 + b2<<12 + b3<<8 + b4<<4
	b1, b2, b3, b4 = p.y&0xf000, p.y&0x0f00, p.y&0x00f0, p.y&0x000f
	sum += b1<<12 + b2<<8 + b3<<4 + b4
	return sum
}

func benchmarkSearch(name string, b *testing.B, search func(graph.Graph, int, int) []int, n int) {
	g := createBenchmarkGraph(n)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = search(g, 1, n)
	}
}

func BenchmarkSearchBFS100(b *testing.B) {
	benchmarkSearch("BFS", b, BFS, 100)
}

func BenchmarkSearchDFS100(b *testing.B) {
	benchmarkSearch("DFS", b, DFS, 100)
}

func BenchmarkSearchDijkstra100(b *testing.B) {
	benchmarkSearch("Dijkstra", b, Dijkstra, 100)
}

func BenchmarkSearchAStar100(b *testing.B) {
	benchmarkSearch("AStar", b, AStar(h), 100)
}

func BenchmarkSearchBFS1000(b *testing.B) {
	benchmarkSearch("BFS", b, BFS, 1000)
}

func BenchmarkSearchDFS1000(b *testing.B) {
	benchmarkSearch("DFS", b, DFS, 1000)
}

func BenchmarkSearchDijkstra1000(b *testing.B) {
	benchmarkSearch("Dijkstra", b, Dijkstra, 1000)
}

func BenchmarkSearchAStar1000(b *testing.B) {
	benchmarkSearch("AStar", b, AStar(h), 1000)
}
