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

func checksum(n *node) bool {
	if n == nil || (n.leftnode == nil && n.rightnode == nil) {
		return true
	}
	var left_sum, right_sum int
	if n.leftnode != nil {
		left_sum = n.leftnode.data
	}
	if n.rightnode != nil {
		right_sum = n.rightnode.data
	}
	if (left_sum+right_sum == n.data) && checksum(n.rightnode) && checksum(n.leftnode) {
		return true
	}
	return false
}

func main() {
	/*
		     10
		   /   \
		  8     2
		 / \   /
		4   4 2
		But the following is not:
		    10
		   / \
		  8   2
		 /  \   \
		 4   4   2
	*/
	var t = tree{}
	t.root = &node{data: 10}
	t.root.leftnode = &node{data: 8}
	t.root.rightnode = &node{data: 2}
	t.root.leftnode.leftnode = &node{data: 4}
	t.root.leftnode.rightnode = &node{data: 4}
	t.root.rightnode.leftnode = &node{data: 5}
	//t.root.rightnode.rightnode = &node{data: 3}
	fmt.Println(checksum(t.root))
}
