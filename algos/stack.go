package main

import "fmt"

type stack struct {
	nodes []node
}

type node struct {
	data int
}

func makeStack() stack {
	return stack{
		nodes: []node{},
	}
}

func makeNode(data int) node {
	return node{
		data: data,
	}
}

// push add an element into the slice
func (s *stack) push(data int) bool {
	s.nodes = append(s.nodes, makeNode(data))
	return true
}

// pop removes an element
func (s *stack) pop() *node {
	l := len(s.nodes)
	if l == 0 {
		return nil
	}
	n := &s.nodes[l-1:][0]
	s.nodes = s.nodes[:l-1]
	return n
}

// see the next element, but not remove from the stack
func (s *stack) top() *node {
	if len(s.nodes) == 0 {
		return nil
	}
	n := s.nodes[len(s.nodes)-1]
	return &n
}

func main() {
	s := makeStack()

	s.push(1)
	s.push(2)
	s.push(3)

	fmt.Println(s.pop().data)
	fmt.Println(s.pop().data)
	fmt.Println(s.pop().data)
}
