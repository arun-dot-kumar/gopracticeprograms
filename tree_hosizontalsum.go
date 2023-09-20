package main

import "fmt"

type Node struct {
	leftnode  *Node
	rightnode *Node
	data      int
}

type tree struct {
	root *Node
}

func findutil(node *Node, h int, a map[int]int) {
	if node == nil {
		return
	}
	findutil(node.leftnode, h+1, a)

	if _, ok := a[h]; ok {
		a[h] = a[h] + node.data
	} else {
		a[h] = node.data
	}
	findutil(node.rightnode, h+1, a)
}

func find(node *Node, h int) {
	var a = make(map[int]int)
	findutil(node, h, a)
	fmt.Println(a)

}

func main() {
	var t = tree{}
	t.root = &Node{data: 10}
	t.root.leftnode = &Node{data: 8}
	t.root.rightnode = &Node{data: 9}
	t.root.leftnode.leftnode = &Node{data: 1}
	t.root.leftnode.rightnode = &Node{data: 2}
	t.root.rightnode.leftnode = &Node{data: 3}
	t.root.rightnode.rightnode = &Node{data: 4}
	find(t.root, 0)

}
