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

func (d *dl) deleten_n_after_every_m(m int, n int) {
	current := d.head
	for current != nil {
		var count int
		for count < m-1 && current != nil {
			count++
			current = current.next
		}
		if current == nil {
			break
		}
		temp := current
		count = 0
		for count <= n && temp != nil {
			count++
			temp = temp.next
		}
		current.next = temp
		current = current.next
	}

}

func main() {
	d := dl{}
	d.insertAtBeg(3)
	d.insertAtBeg(2)
	d.insertAtBeg(1)
	d.insertAtlast(4)
	d.insertAtlast(5)
	d.insertAtlast(6)
	d.insertAtlast(7)
	d.insertAtlast(8)
	d.print()
	fmt.Println()
	d.deleten_n_after_every_m(3, 2)
	d.print()
}
