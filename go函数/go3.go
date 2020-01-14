package main

import "fmt"

type Node struct {
	value int
	LeftNode *Node
	RightNode *Node
}

func createRootNode(v int) *Node {
	return &Node{v,nil, nil}
}

func (n *Node) TraverseFunc(f func(node *Node)) {
	if n == nil {
		return
	}
	n.LeftNode.TraverseFunc(f)
	f(n)
	n.RightNode.TraverseFunc(f)
}

func (n *Node) Traverse() {
	if n == nil {
		return
	}
	n.LeftNode.Traverse()
	fmt.Println(n.value)
	n.RightNode.Traverse()
}

func main() {
	tree := createRootNode(3)
	tree.LeftNode = &Node{4, nil,nil}
	tree.RightNode = &Node{5,nil,nil}

	tree.LeftNode.RightNode = &Node{2, nil,nil}
	tree.RightNode.LeftNode = &Node{6, nil,nil}
	tree.TraverseFunc(func(node *Node) {
		node.value+=10
	})
	tree.Traverse()
}