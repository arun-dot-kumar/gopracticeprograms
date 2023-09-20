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

func find(node *Node) int {
	if node == nil {
		return 0
	}
	lheight := find(node.leftnode)
	rheight := find(node.rightnode)
	if lheight > rheight {
		return lheight + 1
	} else {
		return rheight + 1
	}
}

func main() {
	/* Create the following Binary Tree
	   1
	  /  \
	 2	 3
	/ \	 / \
	4  5 6  7

	*/
	var t = tree{}
	t.root = &Node{data: 10}
	t.root.leftnode = &Node{data: 8}
	t.root.rightnode = &Node{data: 9}
	t.root.leftnode.leftnode = &Node{data: 1}
	t.root.leftnode.rightnode = &Node{data: 2}
	t.root.rightnode.leftnode = &Node{data: 3}
	t.root.rightnode.rightnode = &Node{data: 4}
	fmt.Println(find(t.root))

}
