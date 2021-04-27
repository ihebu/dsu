// Implementation of the Disjoint-Set data structure.
//
// Also called a Union-Find or Merge-Find set,
// is a data structure that stores a collection of disjoint (non-overlapping) sets.
// Equivalently, it stores a partition of a set into disjoint subsets.
// It provides operations for adding new sets,
// merging sets (replacing them by their union),
// and finding a representative member of a set.
// The last operation allows to find out efficiently if any two elements are in the same or different sets.
//
// more at https://en.wikipedia.org/wiki/Disjoint-set_data_structure

package dsu

type node struct {
	value  interface{}
	parent *node
	size   int
}

// DSU is the type used to the Disjoint Set data structure.
// it maps from a value to a node pointer corresponding to the element in the set.
type DSU struct {
	nodes map[interface{}]*node
}

// New returns a pointer to an empty initialized DSU instance.
func New() *DSU {
	return &DSU{map[interface{}]*node{}}
}

// Contains checks if a given element is present in the disjoint set.
func (d *DSU) Contains(x interface{}) bool {
	_, ok := d.nodes[x]

	return ok
}

// Add takes an element as a parameter and inserts it in the disjoint set.
func (d *DSU) Add(x interface{}) bool {
	if d.Contains(x) {
		return false
	}

	d.nodes[x] = &node{value: x, parent: nil, size: 1}
	return true
}

// Find returns the root element that represents the set to which x belongs to.
// If the element doesn't exist in the set, Find returns the nil value.
func (d *DSU) Find(x interface{}) interface{} {
	if !d.Contains(x) {
		return nil
	}

	node := d.nodes[x]

	if node.parent == nil {
		return x
	}

	d.Find(node.parent.value)

	if node.parent.parent != nil {
		node.parent = node.parent.parent
	}

	return node.parent.value
}

// Union replaces the set containing x and the set containing y with their union.
// Union uses Find to determine the roots of the trees containing x and y.
// If the roots are the same of one of the elements doesn't exist in the set,
// there is nothing more to do. and Union returns false
// Otherwise, the two sets get be merged. This is done by either setting
// the parent element of the element with the smaller size to the other parent
// and the return of the function in this case is true
func (d *DSU) Union(x, y interface{}) bool {
	if !d.Contains(x) || !d.Contains(y) {
		return false
	}

	if d.Find(x) == d.Find(y) {
		return false
	}

	nodex := d.nodes[d.Find(x)]
	nodey := d.nodes[d.Find(y)]

	if nodex.size > nodey.size {
		nodey.parent = nodex
		nodex.size += nodey.size
	} else {
		nodex.parent = nodey
		nodey.size += nodex.size
	}

	return true
}
