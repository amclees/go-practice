package graph

type PointerGraph struct {
	nodes []Node
	edges []Edge
	ne, nn int
	size int
}

type Node struct {
	id int64
	val interface{}
}

type Edge struct {
	src, dst int64
	w int64
}

func (g *PointerGraph) AddNode(id int64, val interface{}) {
	if g.nn > 0 {
		for i, node := range g.nodes {
			if node.id == 0 {
				g.nodes[i] = Node{id, val}
				return
			}
		}
	} else {
		g.nodes = append(g.nodes, Node{id, val})
	}
}

func (g *PointerGraph) AddEdge(src, dst, w int64) {
	if g.ne > 0 {
		for i, edge := range g.edges {
			if edge.src == 0 && edge.dst == 0 {
				g.edges[i] = Edge{src, dst, w}
				return
			}
		}
	} else {
		g.edges = append(g.edges, Edge{src, dst, w})
	}
}

func (g *PointerGraph) Node(id int64) (bool, interface{}) {
	for _, node := range g.nodes {
		if node.id == id {
			return true, node.val
		}
	}
	return false, nil
}

func (g *PointerGraph) Edge(src, dst int64) (bool, int64) {
	for _, edge := range g.edges {
		if edge.src == src && edge.dst == dst {
			return true, edge.w
		}
	}
	return false, 0
}

func (g *PointerGraph) RemoveNode(id int64) bool {
	for _, node := range g.nodes {
		if node.id == id {
			node.id = 0
			node.val = nil
			g.nn += 1
			return true
		}
	}
	return false
}

func (g *PointerGraph) RemoveEdge(src, dst int64) bool {
	for _, edge := range g.edges {
		if edge.src == src && edge.dst == dst {
			edge.src = 0
			edge.dst = 0
			edge.w = 0
			g.ne += 1
			return true
		}
	}
	return false
}

func (g *PointerGraph) Nodes() []int64 {
	n := make([]int64, len(g.nodes))
	for i := range n {
		n[i] = g.nodes[i].id
	}
	return n
}

func (g *PointerGraph) Edges(id int64) [][2]int64 {
	n := make([][2]int64, 1)
	for _, edge := range g.edges {
		if edge.src == id {
			n = append(n, [2]int64{edge.dst, edge.w})
		}
	}
	return n
}

func (g *PointerGraph) EdgeCount() int {
	return len(g.edges) - g.ne
}

func (g *PointerGraph) NodeCount() int {
	return len(g.nodes) - g.nn
}
