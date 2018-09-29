package graph

type PointerGraph struct {
	nodes []Node
	edges []Edge
	ne, nn int
	size int
}

type Node struct {
	id int
	val interface{}
}

type Edge struct {
	src, dst int
	w int
}

func (g *PointerGraph) AddNode(id int, val interface{}) {
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

func (g *PointerGraph) AddEdge(src, dst, w int) {
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

func (g *PointerGraph) Node(id int) (bool, interface{}) {
	for _, node := range g.nodes {
		if node.id == id {
			return true, node.val
		}
	}
	return false, nil
}

func (g *PointerGraph) Edge(src, dst int) (bool, int) {
	for _, edge := range g.edges {
		if edge.src == src && edge.dst == dst {
			return true, edge.w
		}
	}
	return false, 0
}

func (g *PointerGraph) RemoveNode(id int) bool {
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

func (g *PointerGraph) RemoveEdge(src, dst int) bool {
	for i, edge := range g.edges {
		if edge.src == src && edge.dst == dst {
			g.edges[i].src = 0
			g.edges[i].dst = 0
			g.edges[i].w = 0
			g.ne += 1
			return true
		}
	}
	return false
}

func (g *PointerGraph) Nodes() []int {
	n := make([]int, len(g.nodes))
	for i := range n {
		n[i] = g.nodes[i].id
	}
	return n
}

func (g *PointerGraph) Edges(id int) [][2]int {
	n := make([][2]int, 0)
	for _, edge := range g.edges {
		if edge.src == id {
			n = append(n, [2]int{edge.dst, edge.w})
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
