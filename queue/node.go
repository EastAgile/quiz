package queue

type Node struct {
	Word   string
	Suffix string
}

func NewNode(word string, suffix string) *Node {
	return &Node{Word: word, Suffix: suffix}
}
