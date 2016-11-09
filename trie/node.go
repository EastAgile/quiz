package trie

type Node struct {
	Value    byte
	Children []*Node
	End      bool
}

func NewNode() *Node {
	return &Node{Children: make([]*Node, 0)}
}

func (node *Node) addChild(child *Node) {
	node.Children = append(node.Children, child)
}

func (node *Node) Find(c byte) *Node {
	for _, child := range node.Children {
		if child.Value == c {
			return child
		}
	}
	return nil
}
