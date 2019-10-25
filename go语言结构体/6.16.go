package main

import "fmt"

// 链表

type Student struct {
	Id int
	Name string
}

type Node struct {
	Student
	Next *Node
}

func (head *Node) Create() *Node {
	head = nil
	return head
}

func (p *Node) PrintLink() {
	for p != nil {
		fmt.Printf("%d, %s\n", p.Id, p.Name)
	}
}

func (newNode *Node) Insert(head *Node) *Node {
	var p0, p1, p2 *Node

}