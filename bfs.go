package main

import "fmt"

type Graph struct {
	Vertices       int
	AdjecencyGraph map[int][]int
}

func addGraph(v int) *Graph {
	return &Graph{
		Vertices:       v,
		AdjecencyGraph: make(map[int][]int),
	}
}

func (g *Graph) addEdge(src int, dest int) {
	g.AdjecencyGraph[src] = append(g.AdjecencyGraph[src], dest)
}

type BfsQueue struct {
	list []int
}

func (b *BfsQueue) Enqueue(node int) {
	b.list = append(b.list, node)
}

func (b *BfsQueue) Dequeue() {
	if b.isEmpty() {
		return
	}
	b.list = b.list[1:]
}

func (b *BfsQueue) isEmpty() bool {
	if len(b.list) == 0 {
		return true
	}
	return false
}

func bfs(g *Graph, node int) {
	visited := make(map[int]bool)
	var bqueue BfsQueue
	bqueue.Enqueue(node)
	visited[node] = true
	fmt.Printf("%d,", node)
	for !bqueue.isEmpty() {
		curNode := bqueue.list[0]
		for _, v := range g.AdjecencyGraph[curNode] {
			if !visited[v] {
				visited[v] = true
				fmt.Printf("%d,", v)
				bqueue.Enqueue(v)
			}
		}
		bqueue.Dequeue()
	}
}

func main() {
	g := addGraph(5)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(2, 5)
	g.addEdge(2, 6)
	bfs(g, 0)
}
