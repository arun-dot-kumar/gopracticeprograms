package main

import "fmt"

type Node struct {
	next *Node
	data int
	prev *Node
}

type Ll struct {
	head *Node
	tail *Node
}

func (l *Ll) insetAtBeg(data int) {
	node := Node{data: data}
	if l.head == nil {
		l.head = &node
		l.tail = &node
		return
	}
	temp := l.head
	l.head.prev = &node
	l.head = &node
	node.next = temp
}

func (l *Ll) insertAtEnd(data int) {
	node := Node{data: data}
	if l.head == nil {
		l.head = &node
		l.tail = &node
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &node
	node.prev = temp
	l.tail = &node
}

func (l *Ll) reverse() {
	current := l.head
	l.head, l.tail = l.tail, l.head
	var node *Node
	for current != nil {
		node = current.next
		current.prev, current.next = current.next, current.prev
		current = node
	}
}
func (l *Ll) printfoward() {
	if l.head != nil {
		temp := l.head
		for temp != nil {
			fmt.Printf("%d->", temp.data)
			temp = temp.next
		}
	}
}

func (l *Ll) printreverse() {
	if l.head != nil {
		temp := l.tail
		for temp != nil {
			fmt.Printf("%d->", temp.data)
			temp = temp.prev
		}
	}
}

func main() {
	l := Ll{}
	l.insetAtBeg(4)
	l.insetAtBeg(2)
	l.insetAtBeg(3)
	l.insetAtBeg(1)
	l.insertAtEnd(5)
	l.insertAtEnd(6)
	l.printfoward()
	fmt.Println()
	l.printreverse()
	fmt.Println()
	fmt.Println("after reverse")
	l.reverse()
	l.printreverse()
	fmt.Println()
	l.printfoward()
}
