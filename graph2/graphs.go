package graph2

import "github.com/ds0nt/cs-mind-grind/disjoint_set"

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

// AddEdge is a convinience method for ug.AddEdge
func (n *Node) AddEdge(ug *UndirectedGraph, b *Node) {
	ug.AddEdge(n, b)
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
}

type WeightedGraph struct {
}
