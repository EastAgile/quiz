package queue

type Queue struct {
	nodes []*Node
}

func NewQueue() *Queue {
	return &Queue{nodes: make([]*Node, 0)}
}

func (q *Queue) Push(n *Node) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) Pop() *Node {
	if len(q.nodes) == 0 {
		return nil
	}
	var n *Node
	n, q.nodes = q.nodes[0], q.nodes[1:]
	return n
}
