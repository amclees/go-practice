package graph

type MatrixGraph struct {
	matrix [][]int
	nodes []interface{}
	nc, ec int
}

func NewMatrixGraph(nodeHint int) MatrixGraph {
	g := MatrixGraph{}
	g.nodes = make([]interface{}, nodeHint)
	g.matrix = make([][]int, nodeHint)
	for i := range g.matrix {
		g.matrix[i] = make([]int, nodeHint)
	}

	g.AddNode(0, interface{}(nil))
	g.nc -= 1

	return g
}

func (g *MatrixGraph) AddNode(id int, val interface{}) {
	if id >= len(g.nodes) {
		for i := 0; i < len(g.nodes) - id; i++ {
			g.nodes = append(g.nodes, interface{}(nil))
		}
		g.nodes = append(g.nodes, val)
		g.nc += 1

		for i, row := range g.matrix {
			g.matrix[i] = append(row, 0)
		}
		g.matrix = append(g.matrix, make([]int, id))
	} else {
		g.nodes[id] = val
	}
}

func (g *MatrixGraph) AddEdge(src, dst, w int) {
	if src >= len(g.nodes) {
		g.AddNode(src, interface{}(nil))
	}
	if dst >= len(g.nodes) {
		g.AddNode(dst, interface{}(nil))
	}

	if g.matrix[src][dst] == 0 {
		g.ec += 1
	}
	g.matrix[src][dst] = w
}

func (g *MatrixGraph) Node(id int) (bool, interface{}) {
	if id >= len(g.nodes) || g.nodes[id] == interface{}(nil) {
		return false, interface{}(nil)
	}
	return true, g.nodes[id]
}

func (g *MatrixGraph) Edge(src, dst int) (bool, int) {
	if src >= len(g.nodes) || dst >= len(g.nodes) {
		return false, 0
	}
	val := g.matrix[src][dst]
	return val != 0, val
}

func (g *MatrixGraph) RemoveNode(id int) bool {
	if id >= len(g.nodes) {
		return false
	}
	g.nodes[id] = interface{}(nil)
	g.nc -= 1
	return true
}

func (g *MatrixGraph) RemoveEdge(src, dst int) bool {
	if src >= len(g.nodes) || dst >= len(g.nodes) {
		return false
	}
	if g.matrix[src][dst] == 0 {
		return false
	}
	g.matrix[src][dst] = 0
	g.ec -= 1
	return true
}

func (g *MatrixGraph) Nodes() []int {
	nodes := make([]int, 0)
	for i, node := range g.nodes {
		if node != interface{}(nil) {
			nodes = append(nodes, i)
		}
	}
	return nodes
}

func (g *MatrixGraph) Edges(id int) [][2]int {
	if id > len(g.nodes) {
		return [][2]int{}
	}

	edges := make([][2]int, 0)
	for i, weight := range g.matrix[id] {
		if weight != 0 {
			edges = append(edges, [2]int{i, weight})
		}
	}
	return edges
}

func (g *MatrixGraph) EdgeCount() int {
	return g.ec
}

func (g *MatrixGraph) NodeCount() int {
	return g.nc
}
