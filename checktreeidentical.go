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

func find(t1 *Node, t2 *Node) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1.data != t2.data {
		return false
	}
	var status bool
	status = find(t1.leftnode, t2.leftnode)
	if !status {
		return false
	}
	status = find(t1.rightnode, t2.rightnode)
	if !status {
		return false
	}
	return true
}

func main() {
	/* Create the following Binary Tree
	Output: Both trees are identical

	Input:             1                    1
	                 /   \                /   \
	               2      3            5      3
	                     /             /
	                  4             4

	Output: Trees are not identical
	*/
	var t = tree{}
	t.root = &Node{data: 1}
	t.root.leftnode = &Node{data: 2}
	t.root.rightnode = &Node{data: 3}
	t.root.rightnode.leftnode = &Node{data: 4}

	var t2 = tree{}
	t2.root = &Node{data: 1}
	t2.root.leftnode = &Node{data: 2}
	t2.root.rightnode = &Node{data: 3}
	t2.root.rightnode.leftnode = &Node{data: 5}
	fmt.Println(find(t.root, t2.root))

}
