package main

import "fmt"

type lNode struct {
	prevNode *lNode
	data     int
	nextNode *lNode
}

type DobleLL struct {
	head *lNode
}

func (d *DobleLL) insertToDLL(data int) {
	if d.head == nil {
		d.head = &lNode{data: data}
	} else {
		cur := d.head
		for cur.nextNode != nil {
			cur = cur.nextNode
		}
		cur.nextNode = &lNode{data: data}
		cur.nextNode.prevNode = cur
	}
}

func (d *DobleLL) printDll() {
	if d.head == nil {
		return
	}
	cur := d.head
	for cur != nil {
		fmt.Printf("%d->", cur.data)
		cur = cur.nextNode
	}
	fmt.Println()
}

func (d *DobleLL) reverseDll() {
	if d.head == nil {
		return
	}
	var prev *lNode
	cur := d.head
	for cur != nil {
		next := cur.nextNode
		cur.nextNode = prev
		prev = cur
		cur = next
	}
	d.head = prev
}

func main() {
	dll := &DobleLL{}
	dll.insertToDLL(1)
	dll.insertToDLL(2)
	dll.insertToDLL(3)
	dll.insertToDLL(4)
	dll.insertToDLL(5)
	dll.printDll()
	dll.reverseDll()
	dll.printDll()
}
