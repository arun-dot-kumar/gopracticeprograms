package main

import (
	"fmt"
)

type Node struct {
	next *Node
	data int
}

type dl struct {
	head *Node
}

func (d *dl) insertAtBeg(data int) {
	node := &Node{data: data}
	if d.head == nil {
		d.head = node
		return
	}
	temp := d.head
	d.head = node
	node.next = temp
}

func (d *dl) insertAtlast(data int) {
	node := &Node{data: data}
	if d.head == nil {
		d.insertAtBeg(data)
		return
	}
	temp := d.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
}

func (d *dl) print() {
	if d.head == nil {
		return
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("%d->", temp.data)
		temp = temp.next
	}
}

func (d *dl) deletekthposition(k int) {
	var count = 1
	var current = d.head
	for count < k-1 && current != nil {
		count++
		current = current.next
	}
	temp := current.next
	current.next = temp.next
}

func swapPairs(head *Node) *Node {
	// Create a dummy node to simplify edge cases
	dummy := &Node{data: 0, next: head}
	prev := dummy

	for head != nil && head.next != nil {
		// Nodes to be swapped
		node1 := head
		node2 := head.next

		// Update the pointers to perform the swap
		prev.next = node2
		node1.next = node2.next
		node2.next = node1

		// Move to the next pair
		prev = node1
		head = node1.next
	}

	return dummy.next
}

func main() {
	d := dl{}
	d.insertAtlast(1)
	d.insertAtlast(2)
	d.insertAtlast(3)
	d.insertAtlast(4)
	d.insertAtlast(5)
	d.insertAtlast(6)
	d.insertAtlast(7)
	d.insertAtlast(8)
	d.insertAtlast(9)
	d.insertAtlast(10)
	d.insertAtlast(11)
	d.insertAtlast(12)
	d.print()
	fmt.Println()
	d.head = swapPairs(d.head)
	d.print()
}
