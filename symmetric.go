package main

import "fmt"

type node struct {
	data      int
	leftnode  *node
	rightnode *node
}
type tree struct {
	root *node
}

func checkmirror(node1 *node, node2 *node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 != nil && node2 != nil {
		if node1.data == node2.data {
			return checkmirror(node1.rightnode, node2.leftnode) && checkmirror(node1.leftnode, node2.rightnode)
		}
	}
	return false
}

func checksymmetric(node *node) {
	if checkmirror(node, node) {
		fmt.Println("symmetric")
	} else {
		fmt.Println("not a symmetic")
	}
}

func main() {
	/*
		     1
		   /   \
		  2     2
		 / \   / \
		3   4 4   3
		But the following is not:
		    1
		   / \
		  2   2
		   \   \
		   3    3
	*/
	var t = tree{}
	t.root = &node{data: 1}
	t.root.leftnode = &node{data: 2}
	t.root.rightnode = &node{data: 2}
	//t.root.leftnode.leftnode = &node{data: 3}
	t.root.leftnode.rightnode = &node{data: 3}
	//t.root.rightnode.leftnode = &node{data: 4}
	t.root.rightnode.rightnode = &node{data: 3}
	checksymmetric(t.root)
}
