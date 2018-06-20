package graph2

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
	var testGraph = NewUndirectedGraph()
	n1, n2, n3 := NewNode(1), NewNode(2), NewNode(3)

	testGraph.AddNode(n1, n2, n3)
	testGraph.AddEdge(n1, n2)
	testGraph.AddEdge(n2, n3)

	assert.False(t, testGraph.IsCyclic())

	testGraph.AddEdge(n1, n3)

	assert.True(t, testGraph.IsCyclic())
}
