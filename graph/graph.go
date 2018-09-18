package graph

type Graph interface {
	AddNode(id int64, val interface{})
	AddEdge(src, dst, w int64)

	Node(id int64) (bool, interface{})
	Edge(src, dst int64) (bool, int64)

	RemoveNode(id int64) bool
	RemoveEdge(src, dst int64) bool

	Nodes() []int64
	Edges(id int64) [][2]int64
	EdgeCount() int
	NodeCount() int
}
