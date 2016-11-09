package queue

// Node of Queue
type Node struct {
	Word   string
	Suffix string
}

// NewNode init Node
func NewNode(word string, suffix string) *Node {
	return &Node{Word: word, Suffix: suffix}
}
