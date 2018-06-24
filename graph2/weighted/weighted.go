package weighted

import (
	"container/heap"
	"math"

	"github.com/ds0nt/cs-mind-grind/graph2/directed"
)

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

type Node struct {
	Value interface{}
	Out   []*Edge
	In    []*Edge
}

type Edge struct {
	Weight float32
	To     *Node
	From   *Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: []*Node{},
		Edges: []*Edge{},
	}
}

func NewNode(val interface{}) *Node {
	return &Node{
		Value: val,
		Out:   []*Edge{},
		In:    []*Edge{},
	}
}

// AddNode adds a node to the graph
func (wg *Graph) AddNode(nodes ...*Node) {
	wg.Nodes = append(wg.Nodes, nodes...)
}

// AddEdge adds an edge between two nodes in the graph
func (wg *Graph) AddEdge(from, to *Node, weight float32) {
	edge := Edge{
		Weight: weight,
		To:     to,
		From:   from,
	}
	wg.Edges = append(wg.Edges, &edge)
	from.Out = append(from.Out, &edge)
	to.In = append(to.In, &edge)
}

// IsCyclic deermines if a weighted graph is cyclic.
// Drop the weights and use the same formula as a directed graph
func (wg *Graph) IsCyclic() bool {
	dg := directed.NewDirectedGraph()
	nodeMap := map[*Node]*directed.Node{}
	for _, n := range wg.Nodes {
		_n := directed.NewNode(n.Value)
		nodeMap[n] = _n
		dg.AddNode(_n)
	}

	for _, e := range wg.Edges {
		dg.AddEdge(nodeMap[e.From], nodeMap[e.To])
	}

	return dg.IsCyclic()
}

// ShortestPath implements dijkstras algorithm
// I'm going to just use a min heap for this..
func (wg *Graph) ShortestPath(source *Node, target *Node) (path []*Node, ok bool) {
	if source == target {
		return []*Node{source}, true
	}
	// initialize map of all nodes with weight of zero
	bestPathWeights := map[*Node]float32{}
	// initialize map to track best path's parent for each node
	bestPathFrom := map[*Node]*Node{}
	for _, n := range wg.Nodes {
		bestPathWeights[n] = math.MaxFloat32
		bestPathFrom[n] = nil
	}
	bestPathWeights[source] = 0

	// create a min heap for our best paths
	pathHeap := WeightedPathHeap{&WeightedPath{Weight: 0, Node: source}}
	heap.Init(&pathHeap)

	for {
		if len(pathHeap) == 0 {
			break
		}
		weightedNode := heap.Pop(&pathHeap).(*WeightedPath)

		for _, edge := range weightedNode.Node.Out {

			weight := weightedNode.Weight + edge.Weight

			if weight < bestPathWeights[edge.To] {
				bestPathWeights[edge.To] = weight
				bestPathFrom[edge.To] = weightedNode.Node
				heap.Push(&pathHeap, &WeightedPath{Weight: weight, Node: edge.To})
			}
		}
	}

	if bestPathFrom[target] == nil {
		return nil, false
	}

	path = []*Node{}
	curr := target
	for curr != source {
		path = append([]*Node{curr}, path...)
		curr = bestPathFrom[curr]
	}
	path = append([]*Node{source}, path...)

	return path, true

}

type WeightedPath struct {
	Weight float32
	Node   *Node
}

// Modified from here: https://golang.org/src/container/heap/example_intheap_test.go
type WeightedPathHeap []*WeightedPath

func (h WeightedPathHeap) Len() int           { return len(h) }
func (h WeightedPathHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h WeightedPathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *WeightedPathHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*WeightedPath))
}

func (h *WeightedPathHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
