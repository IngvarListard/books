package main

import (
	"fmt"
)

var root = new(Node)

type Node struct {
	Value int
	Next  *Node
}

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{Value: v}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists!")
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{Value: v}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(n *Node) {
	if n == nil {
		fmt.Println("EMPTY LIST!")
		return
	}

	for n != nil {
		fmt.Println("Value:", n.Value)
		n = n.Next
	}
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("List is empty")
		return 0
	}

	c := 0
	for t != nil {
		c++
		t = t.Next
	}

	return c
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)
	addNode(root, 100)
	traverse(root)
	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
}
