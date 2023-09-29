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

func merge2ll(d *dl, d1 *dl) dl {
	var res = dl{}
	var count int
	var current = d.head
	var current1 = d1.head
	for current1 != nil && current != nil {
		if count%2 == 0 {
			if res.head == nil {
				res.head = &Node{data: current.data}
			} else {
				res.insertAtlast(current.data)
			}
			current = current.next
		} else {
			if res.head == nil {
				res.head = &Node{data: current1.data}
			} else {
				res.insertAtlast(current1.data)
			}
			current1 = current1.next
		}
		count++
	}
	return res
}

func main() {
	d := dl{}
	d.insertAtlast(1)
	d.insertAtlast(3)
	d.insertAtlast(5)
	d.insertAtlast(7)
	d.insertAtlast(9)
	d.print()
	d1 := dl{}
	d1.insertAtlast(2)
	d1.insertAtlast(4)
	d1.insertAtlast(6)
	d1.insertAtlast(8)
	fmt.Println()
	newll := merge2ll(&d, &d1)
	newll.print()
}
