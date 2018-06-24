package weighted

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	assert.Equal(t, 1, 1)

	var testGraph = NewGraph()
	n1, n2, n3 := NewNode(1), NewNode(2), NewNode(3)

	testGraph.AddNode(n1, n2, n3)
	assert.Len(t, testGraph.Nodes, 3)
	assert.Equal(t, testGraph.Nodes[0].Value, 1)
	assert.Equal(t, testGraph.Nodes[1].Value, 2)
	assert.Equal(t, testGraph.Nodes[2].Value, 3)

	testGraph.AddEdge(n1, n2, 4)
	testGraph.AddEdge(n2, n3, 5)
	assert.Len(t, testGraph.Edges, 2)
	assert.Equal(t, testGraph.Edges[0].From, n1)
	assert.Equal(t, testGraph.Edges[0].To, n2)
	assert.Equal(t, testGraph.Edges[1].From, n2)
	assert.Equal(t, testGraph.Edges[1].To, n3)

	assert.Len(t, n1.In, 0)
	assert.Len(t, n2.In, 1)
	assert.Len(t, n3.In, 1)
	assert.Len(t, n1.Out, 1)
	assert.Len(t, n2.Out, 1)
	assert.Len(t, n3.Out, 0)

	assert.Equal(t, n1.Out[0].From, n1)
	assert.Equal(t, n1.Out[0].To, n2)
	assert.Equal(t, n2.In[0], n1.Out[0])
	assert.Equal(t, n2.Out[0].From, n2)
	assert.Equal(t, n2.Out[0].To, n3)
	assert.Equal(t, n3.In[0], n2.Out[0])
}

// TestIsCyclic is just a test copied from directed graph pkg
// the IsCyclic function of weighted graph just hijacks the directed graphs
// algorithm, but this shows that it works in weighted graph anyways
func TestIsCyclic(t *testing.T) {
	// single edge test, then no startpoint cyclic test
	graph := NewGraph()
	n1, n2 := NewNode(1), NewNode(2)
	graph.AddNode(n1, n2)
	graph.AddEdge(n1, n2, 10)
	assert.False(t, graph.IsCyclic())
	graph.AddEdge(n2, n1, 10)
	assert.True(t, graph.IsCyclic())
	// two edges test, then 1 start point cyclic test
	graph = NewGraph()
	n1, n2, n3 := NewNode(1), NewNode(2), NewNode(3)
	graph.AddNode(n1, n2, n3)
	graph.AddEdge(n1, n2, 10)
	graph.AddEdge(n2, n3, 10)
	assert.False(t, graph.IsCyclic())
	graph.AddEdge(n3, n2, 10)
	assert.True(t, graph.IsCyclic())

	// multi-path acyclic test, multi-tree test, single vertex test, alternate tree cycle test
	graph = NewGraph()
	n1, n2, n3 = NewNode(1), NewNode(2), NewNode(3)
	graph.AddNode(n1, n2, n3)
	graph.AddEdge(n1, n2, 10)
	graph.AddEdge(n1, n3, 10)
	graph.AddEdge(n2, n3, 10)
	graph.AddEdge(n2, n3, 10) // double edge
	assert.False(t, graph.IsCyclic())
	// add three disconnected single vertexes
	n4, n5, n6 := NewNode(4), NewNode(5), NewNode(6)
	graph.AddNode(n4, n5, n6)
	assert.False(t, graph.IsCyclic())
	// connect other vertexes acyclically
	graph.AddEdge(n6, n5, 10)
	graph.AddEdge(n6, n4, 10)
	assert.False(t, graph.IsCyclic())

	// create a cycle
	graph.AddEdge(n4, n6, 10)
	assert.True(t, graph.IsCyclic())

}

func TestHeap(t *testing.T) {
	pathHeap := WeightedPathHeap{
		&WeightedPath{40, nil},
		&WeightedPath{30, nil},
		&WeightedPath{100, nil},
		&WeightedPath{0, nil},
		&WeightedPath{1, nil},
		&WeightedPath{-0.2344, nil},
	}
	heap.Init(&pathHeap)

	assert.Equal(t, heap.Pop(&pathHeap).(*WeightedPath).Weight, float32(-0.2344))
	assert.Equal(t, heap.Pop(&pathHeap).(*WeightedPath).Weight, float32(0))
	assert.Equal(t, heap.Pop(&pathHeap).(*WeightedPath).Weight, float32(1))
	assert.Equal(t, heap.Pop(&pathHeap).(*WeightedPath).Weight, float32(30))
	assert.Equal(t, heap.Pop(&pathHeap).(*WeightedPath).Weight, float32(40))
	assert.Equal(t, heap.Pop(&pathHeap).(*WeightedPath).Weight, float32(100))
}

func TestShortestPath(t *testing.T) {
	graph := NewGraph()
	n1, n2, n3, n4 := NewNode(1), NewNode(2), NewNode(3), NewNode(4)
	n5, n6 := NewNode(5), NewNode(6)
	graph.AddNode(n1, n2, n3, n4, n5, n6)
	graph.AddEdge(n1, n2, 10)
	graph.AddEdge(n1, n3, 2)
	graph.AddEdge(n2, n3, 3)
	graph.AddEdge(n2, n4, 5)
	graph.AddEdge(n3, n4, 6)
	graph.AddEdge(n4, n2, 2)
	graph.AddEdge(n3, n2, 2)
	graph.AddEdge(n4, n1, 1)
	// n5 and n6 are unreachable
	graph.AddEdge(n6, n5, 1)
	graph.AddEdge(n5, n1, 1)

	path, ok := graph.ShortestPath(n1, n2)
	assert.ElementsMatch(t, path, []*Node{n1, n3, n2})
	assert.True(t, ok)
	printPath(path, n1, n2)

	path, ok = graph.ShortestPath(n2, n4)
	assert.ElementsMatch(t, path, []*Node{n2, n4})
	assert.True(t, ok)
	printPath(path, n2, n4)

	path, ok = graph.ShortestPath(n4, n2)
	assert.ElementsMatch(t, path, []*Node{n4, n2})
	assert.True(t, ok)
	printPath(path, n4, n2)

	path, ok = graph.ShortestPath(n2, n1)
	assert.ElementsMatch(t, path, []*Node{n2, n4, n1})
	assert.True(t, ok)
	printPath(path, n2, n1)

	path, ok = graph.ShortestPath(n6, n1)
	assert.ElementsMatch(t, path, []*Node{n6, n5, n1})
	assert.True(t, ok)
	printPath(path, n6, n1)

	path, ok = graph.ShortestPath(n1, n6)
	assert.False(t, ok)
	printPath(path, n6, n1)

	path, ok = graph.ShortestPath(n6, n6)
	assert.True(t, ok)
	assert.ElementsMatch(t, path, []*Node{n6})
	printPath(path, n6, n6)
}

func printPath(path []*Node, s, t *Node) {
	fmt.Printf("Path(%v, %v):\n", s.Value, t.Value)
	if path == nil {
		fmt.Println("no path.")
		return
	}
	for _, n := range path {
		fmt.Print(n.Value, " ")
	}
	fmt.Println()
}
