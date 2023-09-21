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

func bst(arr []int) *node {
	if len(arr) > 0 {
		var mid int
		if len(arr)%2 == 0 {
			mid = (len(arr) / 2) - 1
		} else {
			mid = len(arr) / 2
		}
		var t = node{}
		t.data = arr[mid]
		t.leftnode = bst(arr[:mid])
		t.rightnode = bst(arr[mid+1:])
		return &t
	}
	return nil
}

func preorder(y *node) {
	if y != nil {
		fmt.Printf("%d->", y.data)
		preorder(y.leftnode)
		preorder(y.rightnode)
	}
}

func main() {
	var n = []int{1, 2, 3, 4, 5, 6, 7}
	//var t = tree{}
	//t.root = &node{data: 4}
	//t.root.leftnode = &node{data: 2}
	//t.root.rightnode = &node{data: 5}
	//t.root.leftnode.leftnode = &node{data: 1}
	//t.root.leftnode.rightnode = &node{data: 3}
	//t.root.rightnode.leftnode = &node{data: 5}
	//t.root.rightnode.rightnode = &node{data: 3}
	var y = tree{}
	y.root = bst(n)
	preorder(y.root)
}
