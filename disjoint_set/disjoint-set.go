package disjoint_set

// DisjointSet implements a disjoint set data structure
// https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-046j-design-and-analysis-of-algorithms-spring-2012/lecture-notes/MIT6_046JS12_lec16.pdf
type DisjointSet struct {
	// Items contains all existing items in a hashtable
	// allows for testing item existance in entire disjoint set
	Items map[interface{}]*Set

	// Sets is our forest. the key is the root element
	Sets map[interface{}]*Set
}

func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		Items: map[interface{}]*Set{},
		Sets:  map[interface{}]*Set{},
	}
}

// MakeSet creates a new set containing a single item
func (d *DisjointSet) MakeSet(item interface{}) (set *Set, ok bool) {
	// does the item already exist?
	if _, ok := d.Items[item]; ok {
		return nil, false
	}

	set = &Set{
		[]interface{}{
			item,
		},
	}
	d.Items[item] = set
	d.Sets[item] = set

	return set, true
}

// FindSet returns the representative item for whichever set contains the given item
func (d *DisjointSet) FindSet(item interface{}) (rep interface{}, ok bool) {
	set, ok := d.Items[item]
	if !ok {
		return nil, ok
	}
	rep = set.Items[0]
	return
}

// Union unions two sets in our disjoint set
func (d *DisjointSet) Union(a *Set, b *Set) {
	// TODO: ensure set a and set b exist in our disjoint set

	// delete set B from disjoint set sets
	delete(d.Sets, b.Items[0])

	// update items to point to set A
	for _, v := range b.Items {
		d.Items[v] = a
	}

	// append set B's items into set A
	a.Items = append(a.Items, b.Items...)
}

type Set struct {
	Items []interface{}
}
