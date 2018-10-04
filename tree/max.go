package tree

// Max returns the maximum values
func (tree *Tree) Max() (n *Node) {
	if tree == nil {
		return
	}

	return tree.root.max()
}

func (node *Node) max() (n *Node) {

	if node == nil {
		return
	}

	if node.Right == nil {
		return node
	}

	return node.Right.max()
}
