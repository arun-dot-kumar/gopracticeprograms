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

func (d *dl) reverse() {
	if d.head == nil {
		return
	}
	var current = d.head
	var prev *Node
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	d.head = prev
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

func (d *dl) deleteEveryKthNode(k int) {
	var count = 1
	var current = d.head
	for current != nil {
		if count%k == 0 {
			current.next = current.next.next
		}
		current = current.next
		count++
	}
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
	d.deleteEveryKthNode(2)
	d.print()
}
