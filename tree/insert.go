package tree

// Insert a Node into a tree
func (tree *Tree) Insert(key string) {
	if tree == nil {
		return
	}
	node := &Node{Key: key}
	if tree.root == nil {
		tree.root = node
		return
	}
	tree.root.insert(node)
	return
}
func (node *Node) insert(n *Node) {
	if node == nil || n == nil{
		return
	}
	if node.Key > n.Key {
		if node.Left == nil {
			node.Left = n
		} else {
			node.Left.insert(n)
		}
		return
	}

	if node.Key < n.Key {
		if node.Right == nil {
			node.Right = n
		} else {
			node.Right.insert(n)
		}
	}
	return
}
