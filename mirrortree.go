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

func find(t1 *Node) {
	if t1 == nil {
		return
	}
	find(t1.leftnode)
	find(t1.rightnode)
	temp := t1.leftnode
	t1.leftnode = t1.rightnode
	t1.rightnode = temp
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.leftnode)
		fmt.Printf("%d->", node.data)
		inorder(node.rightnode)
	}

}

func main() {
	/* Create the following Binary Tree
		Output: Both trees are identical

	 	Input:             1                     1
		                 /   \                /     \
		               2      3 - >          3        2
		                     /               \
		                    4                 4

		Output: Trees are not identical
	*/
	var t = tree{}
	t.root = &Node{data: 1}
	t.root.leftnode = &Node{data: 2}
	t.root.rightnode = &Node{data: 3}
	t.root.rightnode.leftnode = &Node{data: 4}
	inorder(t.root)
	find(t.root)
	fmt.Println()
	inorder(t.root)

}
