package main

import (
	"fmt"
)

type Node struct {
	data int
	link *Node
}

type ll struct {
	head *Node
}

func rotateRight(head *ll, k int) *Node {

	var legth int
	var current = head.head
	for current != nil {
		legth++
		current = current.link
	}
	current = head.head
	var count = 1
	for count < legth-k && current.link != nil {
		current = current.link
		count++
	}
	kthnode := current
	for current.link != nil {
		current = current.link
	}
	current.link = head.head
	head.head = kthnode.link
	kthnode.link = nil
	return head.head
}

func print(node *Node) {
	for node != nil {
		fmt.Printf("%d->", node.data)
		node = node.link
	}
}
func main() {
	/*
	   input: 1->2->3->4->5
	   k: 2

	   rotate 1: 5->1->2->3->4
	   rotate 2: 4->5->1->2->3

	*/
	l := ll{head: nil}
	l.head = &Node{data: 10}
	l.head.link = &Node{data: 20}
	l.head.link.link = &Node{data: 30}
	l.head.link.link.link = &Node{data: 40}
	l.head.link.link.link.link = &Node{data: 50}
	l.head.link.link.link.link.link = &Node{data: 60}
	print(l.head)
	fmt.Println()
	//rotateRight(&l, 2)
	l.head = rotateRight(&l, 2)
	print(l.head)
}
