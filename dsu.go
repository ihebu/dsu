package dsu

type node struct {
	value  interface{}
	parent *node
	size   int
}

type DSU struct {
	nodes map[interface{}]*node
}

func New() *DSU {
	return &DSU{map[interface{}]*node{}}
}

func (d *DSU) Contains(x interface{}) bool {
	_, ok := d.nodes[x]

	return ok
}

func (d *DSU) Add(x interface{}) bool {
	if d.Contains(x) {
		return false
	}

	d.nodes[x] = &node{value: x, parent: nil, size: 1}
	return true
}
