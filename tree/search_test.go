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
	t *testing.T
	prefix string
	tr Tree
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

	ok = tr.Delete(7)
	if !ok {
		t.Errorf("Expected deletion of 7 to succeed")
	}

	d, ok = tr.Get(7)
	if ok {
		t.Errorf("Expected tr.Get(7) after deletion to fail with ok = false")
	}

	for _, pair := range pairs {
		if pair[0] == 7 {
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
