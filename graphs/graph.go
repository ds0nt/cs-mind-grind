package main

import "fmt"

func main() {

}

// Graph represents an undirected graph with an adjascency list
type Graph []*Node

// Node is a node in an undirected graph
type Node struct {
	Value     string
	Adjascent []*Node
}

func NewNode(v string) *Node {
	return &Node{v, []*Node{}}
}

func (g *Graph) AddEdge(a, b *Node) {
	a.Adjascent = append(a.Adjascent, b)
	b.Adjascent = append(b.Adjascent, a)
}

// TestGraph makes a graph like this:
/*
	7		   4 -- 3
	|          |
	8 -- 0 -- 1 -- 2
	|
	5 -- 6

*/
func TestGraph() Graph {
	var testGraph = Graph{
		NewNode("0"),
		NewNode("1"),
		NewNode("2"),
		NewNode("3"),
		NewNode("4"),
		NewNode("5"),
		NewNode("6"),
		NewNode("7"),
		NewNode("8"),
	}

	testGraph.AddEdge(testGraph[0], testGraph[1])
	testGraph.AddEdge(testGraph[0], testGraph[8])
	testGraph.AddEdge(testGraph[1], testGraph[2])
	testGraph.AddEdge(testGraph[1], testGraph[4])
	testGraph.AddEdge(testGraph[3], testGraph[4])
	testGraph.AddEdge(testGraph[5], testGraph[8])
	testGraph.AddEdge(testGraph[5], testGraph[6])
	testGraph.AddEdge(testGraph[7], testGraph[8])

	fmt.Print(`
7         4 -- 3
|         |
8 -- 0 -- 1 -- 2
|
5 -- 6
`)

	return testGraph
}

func visitDFS(n *Node, target string, visited map[*Node]struct{}) ([]*Node, bool) {
	if _, ok := visited[n]; ok {
		return nil, false
	}
	visited[n] = struct{}{}

	fmt.Println(".. ", n.Value)

	if target == n.Value {
		return []*Node{n}, true
	}

	for _, adj := range n.Adjascent {
		if path, ok := visitDFS(adj, target, visited); ok {
			return append([]*Node{n}, path...), true
		}
	}
	return nil, false
}

func SearchGraphDFS(s *Node, t string) (path []*Node, ok bool) {
	return visitDFS(s, t, map[*Node]struct{}{})
}
