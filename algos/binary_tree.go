package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (b *BinaryTree) insert(data int) *BinaryTree {
	if b.root == nil {
		b.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		b.root.insert(data)
	}
	return b
}

func (b *BinaryNode) insert(data int) {
	if b == nil {
		return
	}

	if data <= b.data {
		if b.left == nil {
			//if left side is nil, insert in the left
			b.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			//otherwise search recursively until find a nil one
			b.left.insert(data)
		}
	} else {
		if b.right == nil {
			b.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			b.right.insert(data)
		}
	}
}

func main() {
	tree := BinaryTree{}

	tree.
		insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)

	fmt.Println("root", tree.root)
	fmt.Println("left", tree.root.left)
	fmt.Println("right", tree.root.right)

	printTree(os.Stdout, tree.root, 0, 'M')

}

func printTree(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		_, _ = fmt.Fprint(w, " ")
	}
	_, _ = fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	printTree(w, node.left, ns+2, 'L')
	printTree(w, node.right, ns+2, 'R')
}
