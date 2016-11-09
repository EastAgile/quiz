package queue

// Queue struct
type Queue struct {
	nodes []*Node
}

// NewQueue init Queue
func NewQueue() *Queue {
	return &Queue{nodes: make([]*Node, 0)}
}

// Push Add new element to queue
func (q *Queue) Push(n *Node) {
	q.nodes = append(q.nodes, n)
}

// Pop Pop element
func (q *Queue) Pop() *Node {
	if len(q.nodes) == 0 {
		return nil
	}
	var n *Node
	n, q.nodes = q.nodes[0], q.nodes[1:]
	return n
}
