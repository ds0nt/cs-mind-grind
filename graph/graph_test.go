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

func TestDFS3(t *testing.T) {
	g := MakeTestGraph3()

	path, ok := DFS(g[0], "4")
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 0 to 4")
	printPath(path)

	path, ok = DFS(g[4], "6")
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 4 to 6")
	printPath(path)

	_, ok = DFS(g[5], "0")
	if ok {
		t.FailNow()
	}
	fmt.Println("no path from 5 to 0")

	_, ok = DFS(g[6], "2")
	if ok {
		t.FailNow()
	}
	fmt.Println("no path from 6 to 2")
}

func TestBFS3(t *testing.T) {
	g := MakeTestGraph3()

	path, ok := BFS(g[0], "4", len(g))
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 0 to 4")
	printPath(path)

	path, ok = BFS(g[4], "6", len(g))
	if !ok {
		t.FailNow()
	}
	fmt.Println("path from 4 to 6")
	printPath(path)

	_, ok = BFS(g[5], "0", len(g))
	if ok {
		t.FailNow()
	}
	fmt.Println("no path from 5 to 0")

	_, ok = BFS(g[6], "2", len(g))
	if ok {
		t.FailNow()
	}
	fmt.Println("no path from 6 to 2")
}

func TestIsCyclic(t *testing.T) {
	g := MakeTestGraph4()
	fmt.Println("Cyclic:", IsCyclic(g))

	fmt.Println("add edge from 6 to 4")
	g.AddDirectedEdge(g[6], g[4])
	fmt.Println("Cyclic:", IsCyclic(g))

}
