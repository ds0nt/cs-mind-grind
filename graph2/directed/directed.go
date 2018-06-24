package directed

type DirectedGraph struct {
	Nodes []*Node
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		Nodes: []*Node{},
	}
}

// Node is a node in a graph
type Node struct {
	Value     interface{}
	Adjascent []*Node
}

// Edge is an edge between two nodes
type Edge struct {
	A, B *Node
}

func NewNode(value interface{}) *Node {
	return &Node{
		Value:     value,
		Adjascent: []*Node{},
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
// simply do a DFS and if we hit any node twice, it's cyclic.
// -- we create a 'virtual' starting vertex that has outgoing edges to every vertex in the graph
// -- this method allows us to DFS a disconnected graph. Seems like hax, but i tried it on paper ^.^
// -- and it seems relatively easy to prove.. we definitally do not add a cycle by adding a vertex with indegree 0..
// -- and if we start at a vertex with indegree > 0 and there is a cycle involving anything before it, then
// -- we should reach that node from this node anyways, before DFSsing from there from our virtual vertex.
func (dg *DirectedGraph) IsCyclic() bool {
	if len(dg.Nodes) == 0 {
		return false
	}

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
