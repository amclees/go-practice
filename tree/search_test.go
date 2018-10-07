package tree

import (
	"strings"
	"testing"
)

type searchTree interface {
	Get(key int) (interface{}, bool)
	Insert(key int, val interface{})
	Delete(key int) bool
}

type wrappedTest struct {
	t      *testing.T
	prefix string
	tr     Tree
}

func (t *wrappedTest) Errorf(str string, vals ...interface{}) {
	vals = append(vals, String(t.tr))
	t.t.Errorf(strings.Join([]string{t.prefix, str, "\n%v\n"}, ": "), vals...)
}

func testSearchTree(t *wrappedTest, tr searchTree) {
	tr.Insert(0, 1)

	d, ok := tr.Get(0)
	if !ok || d != 1 {
		t.Errorf("Expected initial tr.Get(0) = 1, got %v", d)
	}

	pairs := [][2]int{{20, 2}, {-5, 3}, {7, 4}, {2, 5}, {3, 4}, {30, 8}, {9, 10}, {10, 11}}
	for _, pair := range pairs {
		tr.Insert(pair[0], pair[1])
	}

	toDelete := []int{7, 20, 3}
	deleted := make(map[int]bool)
	for _, n := range toDelete {
		ok = tr.Delete(n)
		if !ok {
			t.Errorf("Expected deletion of %d to succeed", n)
		}
		d, ok = tr.Get(n)
		if ok {
			t.Errorf("Expected tr.Get(%d) after deletion to fail with ok = false", n)
		}
		deleted[n] = true
	}

	for _, pair := range pairs {
		if _, ok = deleted[pair[0]]; ok {
			continue
		}
		d, ok = tr.Get(pair[0])
		if !ok {
			t.Errorf("Expected tr.Get(%d) to return ok", pair[0])
		} else if d != pair[1] {
			t.Errorf("Expected tr.Get(%d) = %d, got %d", pair[0], pair[1], d)
		}
	}
}

func TestBST(t *testing.T) {
	tr := &BST{}
	testSearchTree(&wrappedTest{t, "BST", Tree(tr)}, searchTree(tr))
}

func TestAVL(t *testing.T) {
	tr := &AVL{h: -1}
	testSearchTree(&wrappedTest{t, "AVL", tr}, tr)
}
