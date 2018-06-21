package graph2

import (
	"github.com/ds0nt/cs-mind-grind/disjoint_set"
)

// Node is a node in a graph
type Node struct {
	Value     interface{}
	Adjascent []*Node
}

// Edge is an edge between two nodes
type Edge struct {
	A, B *Node
}

// UndirectedGraph is an undirected graph
type UndirectedGraph struct {
	Nodes []*Node
	Edges []*Edge
}

func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		Edges: []*Edge{},
		Nodes: []*Node{},
	}
}

func NewNode(value interface{}) *Node {
	return &Node{
		Value:     value,
		Adjascent: []*Node{},
	}
}

// AddNode adds a node to the graph
func (ug *UndirectedGraph) AddNode(nodes ...*Node) {
	ug.Nodes = append(ug.Nodes, nodes...)
}

// AddEdge adds an edge between two nodes in the graph
func (ug *UndirectedGraph) AddEdge(a, b *Node) {
	a.Adjascent = append(a.Adjascent, b)
	b.Adjascent = append(b.Adjascent, a)
	ug.Edges = append(ug.Edges, &Edge{a, b})
}

func (ug *UndirectedGraph) IsCyclic() bool {

	djs := disjoint_set.NewDisjointSet()

	for _, n := range ug.Nodes {
		djs.MakeSet(n)
	}

	for _, e := range ug.Edges {
		repA, _ := djs.FindSet(e.A)
		repB, _ := djs.FindSet(e.B)
		if repA == repB {
			return true
		}
		djs.Union(djs.Sets[repA], djs.Sets[repB])
	}

	return false
}

type DirectedGraph struct {
	Nodes []*Node
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		Nodes: []*Node{},
	}
}

// AddNode adds a node to the graph
func (dg *DirectedGraph) AddNode(nodes ...*Node) {
	dg.Nodes = append(dg.Nodes, nodes...)
}

// AddEdge adds an edge between two nodes in the graph
func (dg *DirectedGraph) AddEdge(from, to *Node) {
	from.Adjascent = append(from.Adjascent, to)
}

// IsCyclic detects if the directed graph has a cycle in it
// simply do a DFS and if we hit any node twice, it's cyclic
func (dg *DirectedGraph) IsCyclic() bool {
	if len(dg.Nodes) == 0 {
		return false
	}

	// DETERMINE START NODES IS CURRENTLY FLAWED
	// a cyclic subgraph with all vertexes indegree > 0 will not be searched
	// I better find all subgraphs by doing this: https://stackoverflow.com/a/1348995/2532523
	// caught this in my extensive unit testing ^.^
	// count all indegrees
	// SOLUTION: http://www.cs.yale.edu/homes/aspnes/pinewiki/DepthFirstSearch.html
	// use a virtual node connected to all other vertexes and dfs from there
	// (or just use a for loop on all vertexes)
	// indegrees := map[*Node]int{}
	// for _, n := range dg.Nodes {
	// 	indegrees[n] = 0
	// }

	// for _, n := range dg.Nodes {
	// 	for _, adj := range n.Adjascent {
	// 		indegrees[adj]++
	// 	}
	// }

	// // find nodes with no indegrees
	// startNodes := []*Node{}
	// for n, in := range indegrees {
	// 	if in == 0 {
	// 		startNodes = append(startNodes, n)
	// 	}
	// }

	// // if we have startNodes then it's already cyclic
	// if len(startNodes) == 0 {
	// 	return true
	// }
	// // END FLAWEDNESS

	visited := map[*Node]bool{}
	for _, n := range dg.Nodes {
		if _, ok := visited[n]; ok {
			continue
		}
		if hasBackEdges(n, []*Node{}, visited) {
			return true
		}
	}

	return false
}

func hasBackEdges(n *Node, ancestors []*Node, visited map[*Node]bool) bool {
	visited[n] = true
	// ancestor of self also counts as cycle.. :p
	ancestors = append(ancestors, n)

	for _, adj := range n.Adjascent {
		// have we visited this node?
		if _, ok := visited[adj]; ok {
			// is this node a parent? if yes, we have a backedge
			for _, old := range ancestors {
				if old == adj {
					return true
				}
			}
			// if it's not a parent ignore
			continue
		}

		// if we have not visited it yet, recurse into it
		if hasBackEdges(adj, ancestors, visited) {
			return true
		}
	}
	return false
}

type WeightedGraph struct {
}
