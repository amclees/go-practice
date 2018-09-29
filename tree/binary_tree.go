package tree

type BinaryTree struct {
	key int
	left *BinaryTree
	right *BinaryTree
}

func (t *BinaryTree) K() int {
	return 2
}

func (t *BinaryTree) Root() Node {
	return Node(t)
}

func (t *BinaryTree) Key() int {
	return t.key
}

func (t *BinaryTree) Children() []Node {
	l := Node(nil)
	if (t.left != nil) {
		l = Node(t.left)
	}
	r := Node(nil)
	if (t.right != nil) {
		r = Node(t.right)
	}
	return []Node{l, r}
}
