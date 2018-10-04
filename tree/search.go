package tree

import (
	"errors"
)

// Search returns a copy of the Node in the tree if found
func (tree *Tree) Search(key string) (node Node, err error) {
	if tree == nil {
		err = errors.New("Object not found")
		return
	}

	return tree.root.search(key)
}

func (n *Node) search(key string) (node Node, err error) {
	if n.Key == key {
		node.Key = n.Key
		node.Payload = n.Payload
		return
	}
	if n.Key > key {
		return n.Left.search(key)
	}

	return n.Right.search(key)
}
