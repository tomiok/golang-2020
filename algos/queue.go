package main

import "fmt"

//first in first out
// add put an element in the end
// remove dequeue the first element in the slice
type Queue struct {
	nodes []NodeQ
}

func main() {
	q := makeQueue()

	q.add(1)
	q.add(2)
	q.add(3)

	fmt.Println("size", q.size())

	node := q.remove()
	fmt.Println("data", node.data)
	fmt.Println("size", q.size())
	q.add(4)
	node = q.remove()
	fmt.Println("data", node.data)
	q.add(5)
	fmt.Println("size", q.size())
}

func makeQueue() Queue {
	return Queue{
		nodes: []NodeQ{},
	}
}

type NodeQ struct {
	data int
}

func makeNodeQ(data int) NodeQ {
	return NodeQ{data: data}
}

func (q *Queue) add(data int) bool {
	node := makeNodeQ(data)
	q.nodes = append(q.nodes, node)
	return true
}

func (q *Queue) remove() *NodeQ {
	if len(q.nodes) == 0 {
		return nil
	}
	dequeue := q.nodes[0] //FIFO
	q.nodes = q.nodes[1:]
	return &dequeue
}

func (q *Queue) size() int {
	return len(q.nodes)
}
