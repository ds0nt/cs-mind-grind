package main

type DisjointSet struct {
	Nodes map[*Node]struct{}
}

func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		Nodes: map[*Node]struct{}{},
	}
}

func (s *DisjointSet) Merge(b *DisjointSet) (ok bool) {
	for n := range b.Nodes {
		if !s.Add(n) {
			return false
		}
	}
	return true
}

func (s *DisjointSet) Add(n *Node) (ok bool) {
	_, _ok := s.Nodes[n]
	if _ok {
		return false
	}
	s.Nodes[n] = struct{}{}
	return true
}

func IsCyclic(g Graph) bool {
	sets := map[*Node]*DisjointSet{}

	// initialize set {n} for each node
	for _, node := range g {
		sets[node] = NewDisjointSet()
		sets[node].Add(node)
	}

	for _, node := range g {
		for _, adj := range node.Adjascent {
			if sets[node] == sets[adj] {
				return true
			}
			ok := sets[node].Merge(sets[adj])
			if !ok {
				return true
			}
			sets[adj] = sets[node]
		}
	}

	return false
}
