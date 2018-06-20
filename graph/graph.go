package main

import (
	"fmt"
)

func main() {

}

// Graph represents an undirected graph with an adjascency list
type Graph []*Node

// Node is a node in an undirected graph
type Node struct {
	Value     string
	Adjascent []*Node
	Directed  bool
}

func NewNode(v string) *Node {
	return &Node{v, []*Node{}}
}

func (g *Graph) AddEdge(a, b *Node) {
	a.Adjascent = append(a.Adjascent, b)
	b.Adjascent = append(b.Adjascent, a)
}

func (g *Graph) AddDirectedEdge(a, b *Node) {
	a.Adjascent = append(a.Adjascent, b)
	// b.Adjascent = append(b.Adjascent, a)
}

// MakeTestGraph makes a graph like this:
/*
	7		   4 -- 3
	|          |
	8 -- 0 -- 1 -- 2
	|
	5 -- 6
*/
func MakeTestGraph() Graph {
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

// MakeTestGraph2 make a graph like this:
/*
			   5 -- 6
			   |    |
    7		   4 -- 3
	|
	8 -- 0 -- 1 -- 2
*/
func MakeTestGraph2() Graph {
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
	testGraph.AddEdge(testGraph[3], testGraph[6])
	testGraph.AddEdge(testGraph[3], testGraph[4])
	testGraph.AddEdge(testGraph[5], testGraph[4])
	testGraph.AddEdge(testGraph[5], testGraph[6])
	testGraph.AddEdge(testGraph[7], testGraph[8])

	fmt.Print(`
          5 -- 6
          |    |
7         4 -- 3
|         
8 -- 0 -- 1 -- 2
`)

	return testGraph
}

// MakeTestGraph3 make a directed graph like this:
/*
0  -> 1 <-> 2 <-> 3 <- 4 <-> 5
 \|  |/     |------->>-------|
	6
*/
func MakeTestGraph3() Graph {
	var testGraph = Graph{
		NewNode("0"),
		NewNode("1"),
		NewNode("2"),
		NewNode("3"),
		NewNode("4"),
		NewNode("5"),
		NewNode("6"),
	}

	testGraph.AddDirectedEdge(testGraph[0], testGraph[1])
	testGraph.AddDirectedEdge(testGraph[0], testGraph[6])
	testGraph.AddDirectedEdge(testGraph[1], testGraph[6])
	testGraph.AddEdge(testGraph[1], testGraph[2])
	testGraph.AddEdge(testGraph[2], testGraph[3])
	testGraph.AddDirectedEdge(testGraph[4], testGraph[3])
	testGraph.AddEdge(testGraph[4], testGraph[5])
	testGraph.AddDirectedEdge(testGraph[2], testGraph[5])

	fmt.Print(`
0  -> 1 <-> 2 <-> 3 <- 4 <-> 5
 \|  |/     |------->>-------| 
    6
`)

	return testGraph
}

// MakeTestAcyclic make a directed graph like this:

func MakeTestGraph4() Graph {
	fmt.Print(`
2 <- 1 <- 0
     |
     >> 3 <- 4
        ^
        L 5 -> 6
`)
	var testGraph = Graph{
		NewNode("0"),
		NewNode("1"),
		NewNode("2"),
		NewNode("3"),
		NewNode("4"),
		NewNode("5"),
		NewNode("6"),
	}

	testGraph.AddDirectedEdge(testGraph[0], testGraph[1])
	testGraph.AddDirectedEdge(testGraph[1], testGraph[2])
	testGraph.AddDirectedEdge(testGraph[1], testGraph[3])
	testGraph.AddDirectedEdge(testGraph[4], testGraph[3])
	testGraph.AddDirectedEdge(testGraph[5], testGraph[3])
	testGraph.AddDirectedEdge(testGraph[5], testGraph[6])

	return testGraph
}

func dfs(n *Node, target string, visited map[*Node]struct{}) ([]*Node, bool) {
	if _, ok := visited[n]; ok {
		return nil, false
	}

	visited[n] = struct{}{}
	fmt.Println("visited: ", n.Value)

	if target == n.Value {
		return []*Node{n}, true
	}

	for _, adj := range n.Adjascent {
		if path, ok := dfs(adj, target, visited); ok {
			return append([]*Node{n}, path...), true
		}
	}
	return nil, false
}

// DFS does a depth first search on a graph starting from s, searching for value t
func DFS(s *Node, t string) (path []*Node, ok bool) {
	return dfs(s, t, map[*Node]struct{}{})
}

func BFS(s *Node, t string, bufferSize int) (path []*Node, ok bool) {
	type searchNode struct {
		from   *searchNode
		target *Node
	}

	pathToTarget := func(n *searchNode) []*Node {
		path := []*Node{}
		for {
			path = append([]*Node{n.target}, path...)
			if n.from == nil {
				return path
			}

			n = n.from
		}
	}

	visited := map[*Node]struct{}{}
	nodesToSearch := make(chan *searchNode, bufferSize)
	nodesToSearch <- &searchNode{target: s}

	for {
		select {
		case curr := <-nodesToSearch:
			if _, ok := visited[curr.target]; ok {
				continue
			}
			visited[curr.target] = struct{}{}
			fmt.Println("visited: ", curr.target.Value)
			if curr.target.Value == t {
				return pathToTarget(curr), true
			}
			for _, adj := range curr.target.Adjascent {
				// add unvisited adjascent nodes (could double add)
				if _, ok := visited[adj]; !ok {
					nodesToSearch <- &searchNode{from: curr, target: adj}
				}
			}
		default:
			return nil, false
		}
	}

}
