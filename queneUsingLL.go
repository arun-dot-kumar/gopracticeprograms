package main

import "fmt"

type Node struct {
	next *Node
	data int
}

type dl struct {
	rear  *Node
	front *Node
}

func (d *dl) enqueue(data int) {
	node := &Node{data: data}
	if d.rear == nil {
		d.front = node
		d.rear = node
		return
	}
	current := d.front
	for current.next != nil {
		current = current.next
	}
	current.next = node
	d.rear = node
}

func (d *dl) print() {
	current := d.front
	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}
}

func (d *dl) dequeue() {
	if d.rear == nil {
		return
	}
	d.front = d.front.next
	if d.front == nil {
		d.rear = nil
	}
}

func main() {
	d := dl{}
	d.enqueue(1)
	d.enqueue(2)
	d.enqueue(3)
	d.enqueue(4)
	d.enqueue(5)
	d.enqueue(6)
	d.print()
	d.dequeue()
	d.dequeue()
	d.dequeue()
	d.dequeue()
	d.dequeue()
	d.dequeue()
	d.print()
}
