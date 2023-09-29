package main

import "fmt"

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
		node.next = d.head
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
		if temp.next == d.head {
			break
		}
	}
	temp.next = node
	node.next = d.head
}

func (d *dl) print() {
	if d.head == nil {
		return
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("%d->", temp.data)
		temp = temp.next
		if temp == d.head {
			break
		}

	}
}

func (d *dl) delete(key int) {
	if d.head.data == key {
		temp := d.head
		d.head = d.head.next
		current := d.head
		for current.next != nil {
			current = current.next
			if current.next == temp {
				current.next = d.head
				return
			}
		}
	}
	current := d.head
	for current.next.next != nil {
		if current.next.data == key {
			current.next = current.next.next
		}
		return
	}
	current = current.next

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
	d.print()
	d.delete(1)
	fmt.Println()
	d.print()
}
