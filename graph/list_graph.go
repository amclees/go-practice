package graph

type ListGraph struct {
	nodes []List
	ec, nc int
}

type List struct {
	val interface{}
	list [][2]int
}

func NewListGraph(nodeHint, edgeHint int) ListGraph {
	nodes := make([]List, nodeHint)
	for i := range nodes {
		l := make([][2]int, edgeHint)
		nodes[i] = List{list: l}
	}

	return ListGraph{nodes, 0, 0}
}

func (g *ListGraph) AddNode(id int, val interface{}) {
	if id >= len(g.nodes) {
		k := id - len(g.nodes)
		for i := 0; i < k + 2; i++ {
			g.nodes = append(g.nodes, List{list: make([][2]int, 0)})
		}
	}
	g.nodes[id].val = val
	g.nc += 1
}

func (g *ListGraph) AddEdge(src, dst, w int) {
	l := len(g.nodes)
	if l < src {
		g.AddNode(src, interface{}(nil))
	}

	s := g.nodes[src].list
	for i, pair := range s {
		if pair[0] == dst {
			s[i][1] = w
			g.ec += 1
			return
		}
	}

	g.nodes[src].list = append(g.nodes[src].list, [2]int{dst, w})
	g.ec += 1
}

func (g  ListGraph) Node(id int) (bool, interface{}) {
	if id >= len(g.nodes) || g.nodes[id].val == interface{}(nil) {
		return false, interface{}(nil)
	}
	return true, g.nodes[id].val
}

func (g *ListGraph) Edge(src, dst int) (bool, int) {
	if src >= len(g.nodes) || dst >= len(g.nodes) {
		return false, 0
	}

	for _, pair := range g.nodes[src].list {
		if pair[0] == dst {
			return true, pair[1]
		}
	}

	return false, 0
}

func (g *ListGraph) RemoveNode(id int) bool {
	if id >= len(g.nodes) || g.nodes[id].val == interface{}(nil) {
		return false
	}
	g.nodes[id].val = interface{}(nil)
	g.nodes[id].list = make([][2]int, 0)
	g.nc -= 1
	return true
}

func (g *ListGraph) RemoveEdge(src, dst int) bool {
	if src >= len(g.nodes) || dst >= len(g.nodes) {
		return false
	}

	for i, pair := range g.nodes[src].list {
		if pair[0] == dst {
			g.nodes[src].list[i][0] = 0
			g.nodes[src].list[i][1] = 0
			g.ec -= 1
			return true
		}
	}

	return false
}

func (g *ListGraph) Nodes() []int {
	nodes := make([]int, 0)
	for i, node := range g.nodes {
		if node.val != interface{}(nil) {
			nodes = append(nodes, i)
		}
	}
	return nodes
}

func (g *ListGraph) Edges(id int) [][2]int {
	if id >= len(g.nodes) || g.nodes[id].val == interface{}(nil) {
		return [][2]int{}
	}

	edges := make([][2]int, 0)
	for _, edge := range g.nodes[id].list {
		if edge[0] != 0 {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g *ListGraph) EdgeCount() int {
	return g.ec
}

func (g *ListGraph) NodeCount() int {
	return g.nc
}
