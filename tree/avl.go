package tree

type AVL struct {
	key, h int
	val interface{}
	left, right *AVL
}

func (t *AVL) K() int {
	return 2
}

func (t *AVL) Root() Node {
	return Node(t)
}

func (t *AVL) Key() int {
	return t.key
}

func (t *AVL) Children() []Node {
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

func (tr *AVL) Get(key int) (interface{}, bool) {
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

func (tr *AVL) Insert(key int, val interface{}) {
	if tr.h == -1 {
		tr.h = 0
		tr.key = key
		tr.val = val
		return
	}

	path := make([]*AVL, 0, tr.h + 1)
	for tr != nil {
		path = append(path, tr)
		if key <= tr.key {
			tr = tr.left
		} else {
			tr = tr.right
		}
	}

	end := path[len(path) - 1]
	new := &AVL{key: key, h: 0, val: val}
	if key <= end.key {
		end.left = new
	} else if key > end.key {
		end.right = new
	}

	path = append(path, new)
	rebalance(path)
}

func (tr *AVL) Delete(key int) bool {
	// Delete like normal, pushing ancestors to stack
	// Rebalance on the stack
	return false
}

func (tr *AVL) NodeCount() int {
	c := 1
	if tr.left != nil {
		c += tr.left.NodeCount()
	}
	if tr.right != nil {
		c += tr.right.NodeCount()
	}
	return c
}

func (tr *AVL) Height() int {
	return tr.h + 1
}

func (tr *AVL) leftRotate() {
	r := tr.right
	tr.right = r.left
	n := *tr
	r.left = &n
	*tr = *r

	tr.left.setHeight()
	tr.setHeight()
}

func (tr *AVL) rightRotate() {
	l := tr.left
	tr.left = l.right
	n := *tr
	l.right = &n
	*tr = *l

	tr.right.setHeight()
	tr.setHeight()
}

func (tr *AVL) setHeight() {
	lh, rh := tr.lrHeights()

	if lh > rh {
		tr.h = lh + 1
	} else {
		tr.h = rh + 1
	}
}

func (tr *AVL) lrHeights() (int, int) {
	var l, r int
	if tr.left == nil {
		l = -1
	} else {
		l = tr.left.h
	}
	if tr.right == nil {
		r = -1
	} else {
		r = tr.right.h
	}

	return l, r
}

func leftHeavy(tr *AVL) bool {
	l, r := tr.lrHeights()
	return l > r
}

func rightHeavy(tr *AVL) bool {
	l, r := tr.lrHeights()
	return r > l
}

func rebalance(s []*AVL) {
	for i := len(s) - 1; i >= 0; i-- {
		tr := s[i]
		tr.setHeight()

		if tr.left != nil && rightHeavy(tr.left) {
			tr.left.leftRotate()
		}
		if tr.right != nil && leftHeavy(tr.right) {
			tr.right.rightRotate()
		}
		if leftHeavy(tr) {
			tr.rightRotate()
		}
		if rightHeavy(tr) {
			tr.leftRotate()
		}
	}
}
