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

func (g *Graph) Dfsutil(node int, visited map[int]bool) {
	visited[node] = true
	fmt.Printf("%d,", node)
	for _, v := range g.AdjecencyGraph[node] {
		if !visited[v] {
			g.Dfsutil(v, visited)
		}
	}
}

func (g *Graph) dfs() {
	visited := make(map[int]bool)
	g.Dfsutil(0, visited)
}

func main() {
	g := addGraph(5)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(2, 5)
	g.addEdge(2, 6)
	g.dfs()
}
