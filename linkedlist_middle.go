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

func middle(l *ll) {
	var current = l.head
	var length int
	for current != nil {
		length++
		current = current.link
	}
	current = l.head
	var curlen = 1
	var mid int
	if length%2 != 0 {
		mid = (length / 2) + 1
	} else {
		mid = length / 2
	}
	for current != nil {
		if curlen == mid {
			fmt.Println(current.data)
			break
		}
		curlen++
		current = current.link
	}
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
	//l.head.link.link.link.link.link = &Node{data: 60}
	print(l.head)
	fmt.Println()
	middle(&l)
	//l.head = rotateRight(&l, 2)
	//print(l.head)
}
