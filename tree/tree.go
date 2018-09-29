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

func String(tree Tree) string {
	b := strings.Builder{}
	printNode(&b, tree.Root(), "")
	return b.String()
}

func printNode(b *strings.Builder, node Node, pad string) {
	b.WriteString(pad)
	b.WriteString(strconv.Itoa(node.Key()))
	b.WriteByte('\n')
	for _, n := range node.Children() {
		if n == Node(nil) {
			continue
		}
		printNode(b, n, pad + "    ")
	}
}
