package undirected

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirectedGraph(t *testing.T) {
	assert.Equal(t, 1, 1)

	var testGraph = NewUndirectedGraph()
	n1, n2, n3 := NewNode(1), NewNode(2), NewNode(3)

	testGraph.AddNode(n1, n2, n3)
	assert.Len(t, testGraph.Nodes, 3)
	assert.Equal(t, testGraph.Nodes[0].Value, 1)
	assert.Equal(t, testGraph.Nodes[1].Value, 2)
	assert.Equal(t, testGraph.Nodes[2].Value, 3)

	testGraph.AddEdge(n1, n2)
	testGraph.AddEdge(n2, n3)
	assert.Len(t, testGraph.Edges, 2)
	assert.Equal(t, testGraph.Edges[0].A, n1)
	assert.Equal(t, testGraph.Edges[0].B, n2)
	assert.Equal(t, testGraph.Edges[1].A, n2)
	assert.Equal(t, testGraph.Edges[1].B, n3)
}

func TestIsCyclic(t *testing.T) {

	// basic 3 node connected graph
	var testGraph = NewUndirectedGraph()
	n1, n2, n3 := NewNode(1), NewNode(2), NewNode(3)

	testGraph.AddNode(n1, n2, n3)
	testGraph.AddEdge(n1, n2)
	testGraph.AddEdge(n2, n3)

	assert.False(t, testGraph.IsCyclic())

	// create a cycle
	testGraph.AddEdge(n1, n3)
	assert.True(t, testGraph.IsCyclic())

	// disconnected 4 node graph
	testGraph = NewUndirectedGraph()
	n1, n2, n3, n4 := NewNode(1), NewNode(2), NewNode(3), NewNode(4)
	testGraph.AddNode(n1, n2, n3, n4)
	testGraph.AddEdge(n1, n2)
	assert.False(t, testGraph.IsCyclic())
	testGraph.AddEdge(n3, n4)
	assert.False(t, testGraph.IsCyclic())
	// connect graph
	testGraph.AddEdge(n2, n3)
	assert.False(t, testGraph.IsCyclic())
	// create cycle
	testGraph.AddEdge(n4, n1)
	assert.True(t, testGraph.IsCyclic())

	// graph with double edge
	testGraph = NewUndirectedGraph()
	n1, n2 = NewNode(1), NewNode(2)
	testGraph.AddNode(n1, n2)
	testGraph.AddEdge(n1, n2)
	assert.False(t, testGraph.IsCyclic())
	testGraph.AddEdge(n1, n2)
	assert.True(t, testGraph.IsCyclic())

	// graph with edge to self
	testGraph = NewUndirectedGraph()
	n1, n2 = NewNode(1), NewNode(2)
	testGraph.AddNode(n1, n2)
	assert.False(t, testGraph.IsCyclic())
	testGraph.AddEdge(n1, n1)
	assert.True(t, testGraph.IsCyclic())

}
