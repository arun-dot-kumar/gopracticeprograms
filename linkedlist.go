package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	data int
	link *Node
}
type linkedl struct {
	head *Node
	len  int
}

func (ll *linkedl) insertatBeg(data int) {
	var node = Node{data: data}
	if ll.head == nil {
		ll.head = &node
		ll.len += 1
		return
	}
	var temp = ll.head
	ll.head = &node
	ll.head.link = temp
	ll.len += 1

}
func (ll *linkedl) printLL() {
	var llcontent []string
	if ll.head != nil {
		var temp = ll.head
		for temp != nil {
			llcontent = append(llcontent, strconv.Itoa(temp.data))
			temp = temp.link
		}
		fmt.Println(strings.Join(llcontent[:], "->"))
	} else {
		fmt.Println("linked list is empty")
	}

}
func (ll *linkedl) insertatEnd(data int) {
	var node = Node{data: data}
	if ll.head == nil {
		ll.insertatBeg(data)
		ll.len += 1
		return
	}
	var temp = ll.head
	for temp.link != nil {
		temp = temp.link
	}
	temp.link = &node
	ll.len += 1
}

func (d *ll) insertt3rd(position int, data int) {
	node := Node{data: data}
	var count = 1
	var current = d.head
	for count < position-1 && current != nil {
		count++
		current = current.next
	}
	temp := current.next
	current.next = &node
	node.next = temp
}


func (ll *linkedl) reverse() {
	if ll.head == nil {
		return
	}
	var current = ll.head
	var prev *Node
	var node *Node
	for current != nil {
		node = current.link
		current.link = prev
		prev = current
		current = node
	}
	ll.head = prev
}

func main() {
	var ll = linkedl{}
	ll.insertatBeg(70)
	ll.insertatBeg(60)
	ll.insertatBeg(50)
	ll.printLL()
	ll.insertatEnd(80)
	ll.insertatEnd(90)
	ll.insertatEnd(10)
	ll.insertt3rd(3, 3)
	ll.printLL()
	ll.reverse()
	ll.printLL()
}
