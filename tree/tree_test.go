package tree

import (
	"strings"
	"testing"
)

func TestTreePrint(t *testing.T) {
	bt := BinaryTree{key: 31}
	s := make([]BinaryTree, 5)
	for i := range s {
		s[i] = BinaryTree{key: i}
	}
	bt.left = &s[0]
	bt.right = &s[1]
	bt.left.left = &s[2]
	bt.right.left = &s[3]
	bt.right.right = &s[4]

	tree := Tree(&bt)
	str := String(tree)
	estr := []string{"31", "0", "2", "1", "3", "4"}
	for i, l := range strings.Split(strings.TrimSpace(str), "\n") {
		if strings.TrimSpace(l) != estr[i] {
			t.Errorf("Got tree line print %v, expected %v", l, estr[i])
		}
	}
}
