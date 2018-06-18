package main

import (
	"fmt"
	"strings"
	"testing"
)

func printPath(p []*Node) {
	var strs = []string{}
	for _, n := range p {
		strs = append(strs, n.Value)
	}
	fmt.Println(strings.Join(strs, " - "))
}

func TestDFS(t *testing.T) {
	graph := MakeTestGraph()
	path, ok := DFS(graph[0], "7")
	if !ok {
		t.FailNow()
	}
	printPath(path)
}

func TestBFS(t *testing.T) {
	graph := MakeTestGraph()
	path, ok := BFS(graph[0], "7", len(graph))
	if !ok {
		t.FailNow()
	}
	printPath(path)
}

func TestDFS2(t *testing.T) {
	graph := MakeTestGraph2()
	path, ok := DFS(graph[0], "7")
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 0 to 7")
	printPath(path)

	_, ok = DFS(graph[4], "7")
	if ok {
		t.FailNow()
	}
	fmt.Println("no path from 4 to 7")

	path, ok = DFS(graph[4], "6")
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 4 to 6")
	printPath(path)

	path, ok = DFS(graph[4], "4")
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 4 to 4")
	printPath(path)

}

func TestBFS2(t *testing.T) {
	graph := MakeTestGraph2()
	path, ok := BFS(graph[0], "7", len(graph))
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 0 to 7")
	printPath(path)

	_, ok = BFS(graph[4], "7", len(graph))
	if ok {
		t.FailNow()
	}
	fmt.Println("no path from 4 to 7")

	path, ok = BFS(graph[4], "6", len(graph))
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 4 to 6")
	printPath(path)

	path, ok = BFS(graph[4], "4", len(graph))
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 4 to 4")
	printPath(path)

}
