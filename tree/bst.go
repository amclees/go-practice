package tree

type BST struct {
	key         int
	val         interface{}
	left, right *BST
}

func (t *BST) K() int {
	return 2
}

func (t *BST) Root() Node {
	return Node(t)
}

func (t *BST) Key() int {
	return t.key
}

func (t *BST) Children() []Node {
	l := Node(nil)
	if t.left != nil {
		l = Node(t.left)
	}
	r := Node(nil)
	if t.right != nil {
		r = Node(t.right)
	}
	return []Node{l, r}
}

func (tr *BST) Get(key int) (interface{}, bool) {
	if key == tr.key && tr.val != nil {
		return tr.val, true
	} else if key < tr.key {
		if tr.left == nil {
			return nil, false
		}
		return tr.left.Get(key)
	} else {
		if tr.right == nil {
			return nil, false
		}
		return tr.right.Get(key)
	}
}

func (tr *BST) Insert(key int, val interface{}) {
	if key == tr.key && tr.val == nil && tr.left == nil && tr.right == nil {
		// Edge case, root node has 0 key and value nil
		tr.val = val
	} else if key <= tr.key {
		if tr.left == nil {
			tr.left = &BST{key: key, val: val}
			return
		}
		tr.left.Insert(key, val)
	} else {
		if tr.right == nil {
			tr.right = &BST{key: key, val: val}
			return
		}
		tr.right.Insert(key, val)
	}
}

func (tr *BST) Delete(key int) bool {
	if key == tr.key && tr.val != nil {
		if tr.left == nil && tr.right != nil {
			(*tr) = *(tr.right)
			return true
		} else if tr.right == nil && tr.left != nil {
			(*tr) = *(tr.left)
			return true
		} else if tr.right == nil && tr.left == nil {
			return false
		} else {
			post := false
			next, nextParent := tr.pre()
			if next == tr {
				post = true
				next, nextParent = tr.post()
			}

			var nextChild *BST
			if next.left != nil {
				nextChild = next.left
			}
			if next.right != nil {
				nextChild = next.right
			}

			if nextParent != nil && nextChild != nil {
				if post {
					nextParent.right = nextChild
				} else {
					nextParent.left = nextChild
				}
			}

			l, r := tr.left, tr.right
			(*tr) = *next
			tr.left, tr.right = l, r

			return true
		}
	} else if key < tr.key {
		if tr.left == nil {
			return false
		}
		return tr.left.Delete(key)
	} else {
		if tr.right == nil {
			return false
		}
		return tr.right.Delete(key)
	}
}

func (tr *BST) pre() (*BST, *BST) {
	if tr.left == nil {
		return tr, nil
	}

	var node *BST
	prev := tr
	for node = tr.left; node.right != nil; {
		prev = node
		node = node.right
	}

	return node, prev
}

func (tr *BST) post() (*BST, *BST) {
	if tr.right == nil {
		return tr, nil
	}

	var node *BST
	prev := tr
	for node = tr.right; node.left != nil; {
		prev = node
		node = node.left
	}

	return node, prev
}
