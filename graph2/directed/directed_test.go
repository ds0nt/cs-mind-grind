package directed

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectedGraph(t *testing.T) {
	graph := NewDirectedGraph()
	n1, n2 := NewNode(1), NewNode(2)
	graph.AddNode(n1, n2)

	assert.Len(t, graph.Nodes, 2)
	assert.Equal(t, graph.Nodes[0], n1)
	assert.Equal(t, graph.Nodes[1], n2)

	graph.AddEdge(n1, n2)
	assert.Len(t, n1.Adjascent, 1)
	assert.Len(t, n2.Adjascent, 0)
	assert.Equal(t, n1.Adjascent[0], n2)

	graph.AddEdge(n1, n2)
	assert.Len(t, n1.Adjascent, 2)

	graph.AddEdge(n2, n1)
	assert.Len(t, n1.Adjascent, 2)
	assert.Len(t, n2.Adjascent, 1)
	assert.Equal(t, n2.Adjascent[0], n1)

}

func TestDirectedIsCyclic(t *testing.T) {

	// single edge test, then no startpoint cyclic test
	graph := NewDirectedGraph()
	n1, n2 := NewNode(1), NewNode(2)
	graph.AddNode(n1, n2)
	graph.AddEdge(n1, n2)
	assert.False(t, graph.IsCyclic())
	graph.AddEdge(n2, n1)
	assert.True(t, graph.IsCyclic())

	// two edges test, then 1 start point cyclic test
	graph = NewDirectedGraph()
	n1, n2, n3 := NewNode(1), NewNode(2), NewNode(3)
	graph.AddNode(n1, n2, n3)
	graph.AddEdge(n1, n2)
	graph.AddEdge(n2, n3)
	assert.False(t, graph.IsCyclic())
	graph.AddEdge(n3, n2)
	assert.True(t, graph.IsCyclic())

	// multi-path acyclic test, multi-tree test, single vertex test, alternate tree cycle test
	graph = NewDirectedGraph()
	n1, n2, n3 = NewNode(1), NewNode(2), NewNode(3)
	graph.AddNode(n1, n2, n3)
	graph.AddEdge(n1, n2)
	graph.AddEdge(n1, n3)
	graph.AddEdge(n2, n3)
	graph.AddEdge(n2, n3) // double edge
	assert.False(t, graph.IsCyclic())
	// add three disconnected single vertexes
	n4, n5, n6 := NewNode(4), NewNode(5), NewNode(6)
	graph.AddNode(n4, n5, n6)
	assert.False(t, graph.IsCyclic())
	// connect other vertexes acyclically
	graph.AddEdge(n6, n5)
	graph.AddEdge(n6, n4)
	assert.False(t, graph.IsCyclic())

	// create a cycle
	graph.AddEdge(n4, n6)
	assert.True(t, graph.IsCyclic())

}

func TestTopologicalSort(t *testing.T) {
	graph := NewDirectedGraph()
	n1, n2, n3, n4, n5, n6, n7 := NewNode(1), NewNode(2), NewNode(3), NewNode(4), NewNode(5), NewNode(6), NewNode(7)
	graph.AddNode(n6, n7, n1, n2, n3, n4, n5)
	graph.AddEdge(n1, n2)
	graph.AddEdge(n1, n3)
	graph.AddEdge(n2, n3)
	graph.AddEdge(n2, n3) // double edge
	graph.AddEdge(n4, n3)
	graph.AddEdge(n5, n1)
	graph.AddEdge(n7, n6)

	assert.False(t, graph.IsCyclic())

	sorted := graph.TopologicalSort()
	for _, n := range sorted {
		fmt.Println(n.Value)
	}
	// there are multiple valid orders, but the algo returns this one :p
	assert.Equal(t, sorted[0], n7)
	assert.Equal(t, sorted[1], n4)
	assert.Equal(t, sorted[2], n5)
	assert.Equal(t, sorted[3], n6)
	assert.Equal(t, sorted[4], n1)
	assert.Equal(t, sorted[5], n2)
	assert.Equal(t, sorted[6], n3)
}
