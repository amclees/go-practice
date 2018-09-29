package graph

type Graph interface {
	AddNode(id int, val interface{})
	AddEdge(src, dst, w int)

	Node(id int) (bool, interface{})
	Edge(src, dst int) (bool, int)

	RemoveNode(id int) bool
	RemoveEdge(src, dst int) bool

	Nodes() []int
	Edges(id int) [][2]int
	EdgeCount() int
	NodeCount() int
}
