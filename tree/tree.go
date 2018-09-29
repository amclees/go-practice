package tree

import (
	"strconv"
	"strings"
)

type Node interface {
	Key() int
	Children() []Node
}

type Tree interface {
	K() int
	Root() Node
}

type DepthNode struct {
	node  Node
	depth int
}

type printNode struct {
	node  Node
	depth int
	pos   int
}

func String(tree Tree) string {
	// nw represents the width of each node in the string (shorter digits will be padded)
	nw := 1
	// rc is the depth of the tree
	rc := 0

	stack := []DepthNode{DepthNode{tree.Root(), 0}}
	count := 0
	for current := stack[0]; len(stack) != 0; current = stack[len(stack)-1] {
		nd := current.depth + 1
		for _, child := range current.node.Children() {
			if child == Node(nil) {
				continue
			}
			stack = append(stack, DepthNode{child, nd})
		}

		d := digits(current.node.Key())
		if d > nw {
			nw = d
		}

		if current.depth > rc {
			rc = current.depth
		}

		stack = stack[:len(stack)-1]
		count += 1
	}

	// Queue of printNodes from bottom to top, left to right
	ch := make(chan printNode, count)
	writeToChannel(printNode{tree.Root(), 0, 0}, ch, tree.K()/2)

	b := strings.Builder{}
	depth := 0
	// ns is the spacing between nodes at depth depth.
	ns := nw
	for pn := range ch {
		if pn.depth > depth {
			b.WriteByte('\n')
			b.WriteString(strings.Repeat(" ", (rc-1-pn.depth)*nw))
			ns += (rc - 1 - (pn.depth - 1)) + nw
		}
		b.WriteString(padVal(pn.node.Key(), nw))
		b.WriteString(strings.Repeat(" ", ns))
	}

	return b.String()
}

// writeToChannel writes a tree's printNodes to a channel postorder.
func writeToChannel(pn printNode, ch chan printNode, km int) {
	for i, child := range pn.node.Children() {
		writeToChannel(printNode{child, pn.depth + 1, pn.pos + (km - i)}, ch, km)
	}
	ch <- pn
}

func digits(k int) int {
	d := 0
	for ; k > 0; k /= 10 {
		d++
	}
	return d
}

func padVal(val, nw int) string {
	base := strconv.Itoa(val)
	return strings.Repeat(" ", nw-len(base)) + base
}
