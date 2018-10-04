package tree

import "fmt"

// Delete a Node from the Tree
func (tree *Tree) Delete(key string) {
	if tree == nil {
		return
	}

	tree.root = tree.root.delete(key)
	return
}

func (node *Node) delete(key string) (n *Node) {
	if node == nil {
		return
	}

	t := node
	if t.Key == key {
		fmt.Printf("Removing node: %v\n", t.Key)

		if t.Left != nil {
			fmt.Printf("Inserting <right> into <left>, returng <left>\n")
			t.Left.insert(t.Right)
			return t.Left
		}

		return t.Right
	} else if t.Key > key {
		t.Left = t.Left.delete(key)
	} else {
		t.Right = t.Right.delete(key)
	}

	return t
}
