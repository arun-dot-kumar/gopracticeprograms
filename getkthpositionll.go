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

func (d *dl) getatkthposition(k int) {
	var count int
	var length int
	var current = d.head
	for current != nil {
		length++
		current = current.next
	}
	current = d.head
	for count < (length-k) && current != nil {
		count++
		current = current.next
	}
	fmt.Println(current.data)

}

func main() {
	d := dl{}
	d.insertAtBeg(4)
	d.insertAtBeg(3)
	d.insertAtBeg(2)
	d.insertAtBeg(1)
	d.insertAtlast(5)
	d.insertAtlast(6)
	//d.print()
	d.getatkthposition(3)
}
