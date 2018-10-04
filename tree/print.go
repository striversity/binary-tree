package tree

import (
	"fmt"
	"strings"
)

func (tree *Tree) String() string {
	if tree == nil {
		return "<nil>"
	}
	var sb strings.Builder
	p := func(n *Node) {
		sb.WriteString(fmt.Sprintf("%v\n", n))
	}
	tree.root.traverse(p)
	return sb.String()
}
func (node *Node) String() string {
	if node == nil {
		return "<nil>\n"
	}
	var l, r = "<nil>", "<nil>"
	k := node.Key
	if node.Left != nil {
		l = node.Left.Key
	}
	if node.Right != nil {
		r = node.Right.Key
	}

	return fmt.Sprintf("%s -> L:%s\n%s -> R:%s\n", k, l, k, r)
}
