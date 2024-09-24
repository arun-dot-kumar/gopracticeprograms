package main

import "fmt"

package main

import "fmt"

type Node struct {
	data int
	link *Node
}

type LL struct {
	head   *Node
	length int
}

func (n *LL) insertToLL(data int) {
	if n.head == nil {
		n.head = &Node{data: data}
		n.length++
	} else {
		l := n.head
		for l.link != nil {
			l = l.link
		}
		l.link = &Node{data: data}
		n.length++
	}
}

func (n *LL) PrintLL() {
	if n.head == nil {
		return
	} else {
		l := n.head
		for l != nil {
			fmt.Printf("%d->", l.data)
			l = l.link
		}
	}
	fmt.Println()
}

func (l *LL) findMiddleLength() int {
	midPos := (l.length / 2) + 1
	var count int
	if l.head == nil {
		return 0
	} else {
		l := l.head
		for l != nil {
			count++
			if count == midPos {
				return l.data
			}

			l = l.link
		}
	}
	return 0
}



func main() {
	ll := new(LL)
	ll.head = &Node{data: 1}
	ll.insertToLL(2)
	ll.insertToLL(3)
	ll.insertToLL(4)
	ll.insertToLL(5)
	ll.PrintLL()
	fmt.Println("middle element is", ll.findMiddleLength())
}
