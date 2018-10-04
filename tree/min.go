package tree

// Min returns the minimum values
func (tree *Tree) Min() (n *Node) {
	if tree == nil {
		return
	}

	return tree.root.min()
}

func (node *Node) min() (n *Node) {

	if node == nil {
		return
	}

	if node.Left == nil {
		return node
	}

	return node.Left.min()
}
