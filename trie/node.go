package trie

// Node of Trie
type Node struct {
	Value    byte
	Children []*Node
	End      bool
}

// NewNode init
func NewNode() *Node {
	return &Node{Children: make([]*Node, 0)}
}

// Add a child node to children list
func (node *Node) addChild(child *Node) {
	node.Children = append(node.Children, child)
}

// Find a node in children nodes
func (node *Node) Find(c byte) *Node {
	for _, child := range node.Children {
		if child.Value == c {
			return child
		}
	}
	return nil
}
