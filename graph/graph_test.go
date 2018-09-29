package graph

import "testing"

func testGraph(t *testing.T, graph Graph) {
	testCounts(t, graph, 0, 0)

	for i := 1; i <= 10; i++ {
		graph.AddNode(int64(i), interface{}(int64(i)))
	}

	testCounts(t, graph, 10, 0)

	for i := 1; i <= 10; i++ {
		_, d := graph.Node(int64(i))
		if d != int64(i) {
			t.Errorf("Expected node %d to hold %d, was %d", i, i, d)
		}
	}

	for i := 1; i <= 10; i++ {
		edges := graph.Edges(int64(i))
		if len(edges) > 0 {
			t.Errorf("Expected empty edges initially for %d, was %v", i, edges)
		}
	}

	graph.RemoveNode(int64(10))
	graph.RemoveNode(int64(5))

	testCounts(t, graph, 8, 0)

	graph.AddEdge(int64(1), int64(2), int64(5))
	graph.AddEdge(int64(2), int64(3), int64(10))

	testCounts(t, graph, 8, 2)

	edges := graph.Edges(int64(2))
	if len(edges) != 1 {
		t.Errorf("Expected node 2 to have 1 edge, had %d edges", len(edges))
	}
	if edges[0][1] != 10 {
		t.Errorf("Expected node 2's edge to have weight 10, was %d", edges[0][1])
	}

	ok, w := graph.Edge(1, 2)
	if !ok {
		t.Errorf("Expected graph to have edge from node 1 to node 2 prior to removal")
	}
	if w != int64(5) {
		t.Errorf("Expected edge from node 1 to node 2 to have weight 5, was %d", w)
	}

	graph.RemoveEdge(int64(1), int64(2))

	testCounts(t, graph, 8, 1)

	edges = graph.Edges(int64(1))
	if len(edges) != 0 {
		t.Errorf("Expected node 1 to have no edges after its edge was removed, had %v", edges)
	}

	ok, edge := graph.Edge(1, 2)
	if ok {
		t.Errorf("Expected edge from node 1 to node 2 not to exist after removal, but it existed (was %v)", edge)
	}
}

func testCounts(t *testing.T, graph Graph, nodes, edges int) {
	n := graph.NodeCount()
	if n != nodes {
		t.Errorf("Expected node count %d, was %d", nodes, n)
	}
	e := graph.EdgeCount()
	if e != edges {
		t.Errorf("Expected edge count %d, was %d", edges, e)
	}
}

func TestPointerGraph(t *testing.T) {
	graph := PointerGraph{}
	testGraph(t, Graph(&graph))
}
