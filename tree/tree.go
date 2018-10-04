package tree

type (
	// A Tree is a abstration for a binary tree
	Tree struct {
		root *Node
	}
	// Node is a node in a binary tree
	Node struct {
		Key     string
		Payload int // data being managed in the tree
		Left    *Node
		Right   *Node
	}
)

// New creates a new Tree
func New() *Tree {
	return &Tree{}
}
