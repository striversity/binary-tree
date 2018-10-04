package tree

func (tree *Tree) traverse(f func(n *Node)) {
	if tree == nil || f == nil {
		return
	}

	if tree.root == nil {
		return
	}
	tree.root.traverse(f)
}

func (node*Node)traverse(f func(n*Node)){
	if node == nil{
		return
	}

	node.Left.traverse(f)
	f(node)
	node.Right.traverse(f)
}