package graph

type Graph interface {
	AddNode(id int64, val interface{})
	AddEdge(src, dst int64) bool

	Node(id int64) interface{}
	HasEdge(src, dst int64) bool

	RemoveNode(id int64) bool
	RemoveEdge(src, dst int64) bool

	Nodes() []int64
	Edges(id int64) []int64
	Size() int64
}
