package main

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	graph := TestGraph()
	path, ok := SearchGraphDFS(graph[0], "7")
	if !ok {
		t.FailNow()
	}
	for _, n := range path {
		fmt.Print(n.Value, " - ")
	}
	fmt.Println()
}
