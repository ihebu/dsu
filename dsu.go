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

// Add adds a new element. This element is placed into a new set containing only the new element, and the new set is added to the data structure.
// If the element already exists in the data structure, then nothing is done.
func (d *DSU) Add(x interface{}) {
	if !d.Contains(x) {
		d.nodes[x] = &node{value: x, parent: nil, size: 1}
	}
}

// Find returns the root element that represents the set to which x belongs to.
// If the element does not exist in the data structure, it returns the nil value.
func (d *DSU) Find(x interface{}) interface{} {
	if !d.Contains(x) {
		return nil
	}

	node := d.nodes[x]
	root := node

	for root.parent != nil {
		root = root.parent
	}

	for node.parent != nil {
		node.parent, node = root, node.parent
	}

	return root.value
}

// Union replaces the set containing x and the set containing y with their union. It first determines the roots of the sets containing x and y.
// If the roots are the same or one of the elements does not exist in the data structure, there is nothing more to do. Otherwise, the two sets get be merged.
//
// The root of the new set is the root of the set with bigger size. In case both sets have the same size, the root is the root of set y.
//
// It returns true if the merge happened
func (d *DSU) Union(x, y interface{}) bool {
	rootx := d.Find(x)
	rooty := d.Find(y)

	if rootx == nil || rooty == nil || rootx == rooty {
		return false
	}

	nodex := d.nodes[rootx]
	nodey := d.nodes[rooty]

	if nodex.size <= nodey.size {
		nodex, nodey = nodey, nodex
	}

	nodey.parent = nodex
	nodex.size += nodey.size

	return true
}
