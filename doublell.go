package main

import (
	"fmt"
)

type Node struct {
	next *Node
	data int
	prev *Node
}

type dl struct {
	head *Node
	tail *Node
}

func (d *dl) insertAtBeg(data int) {
	node := &Node{data: data}
	if d.head == nil {
		d.tail = node
		d.head = node
		return
	}
	temp := d.head
	d.head = node
	node.next = temp
	temp.prev = node
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
	node.prev = temp
	d.tail = node
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

func (d *dl) reverse() {
	if d.head == nil {
		return
	}
	var current = d.head
	var prev *Node
	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	temp := d.head
	d.head = prev
	d.tail = temp
}

func main() {
	d := dl{}
	d.insertAtBeg(4)
	d.insertAtBeg(3)
	d.insertAtBeg(2)
	d.insertAtBeg(1)
	d.insertAtlast(5)
	d.insertAtlast(6)
	d.print()
	fmt.Println()
	d.reverse()
	d.print()

}
