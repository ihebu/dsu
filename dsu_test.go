package dsu

import (
	"testing"
)

func assertEqual(t *testing.T, got, expected interface{}) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestContains(t *testing.T) {
	d := New()
	node1 := &node{value: 1, parent: nil, size: 1}
	node2 := &node{value: "foo", parent: nil, size: 1}

	d.nodes[1] = node1
	d.nodes["foo"] = node2

	assertEqual(t, d.Contains(1), true)
	assertEqual(t, d.Contains("foo"), true)
	assertEqual(t, d.Contains(0), false)
}

func TestAdd(t *testing.T) {

	t.Run("Adding a new element to the set", func(t *testing.T) {
		d := New()
		assertEqual(t, d.Add(1), true)
	})

	t.Run("Adding an existing element to the set", func(t *testing.T) {
		d := New()
		d.Add(1)

		assertEqual(t, d.Add(1), false)
	})
}

func TestFind(t *testing.T) {
	t.Run("Finding a non existing element", func(t *testing.T) {
		d := New()
		assertEqual(t, d.Find(1), nil)
	})

	t.Run("Finding an existing element", func(t *testing.T) {
		d := New()
		node1 := &node{value: 1, parent: nil}
		node2 := &node{value: 2, parent: nil}
		node3 := &node{value: 3, parent: node2}
		node4 := &node{value: 4, parent: node3}

		d.nodes[1] = node1
		d.nodes[2] = node2
		d.nodes[3] = node3
		d.nodes[4] = node4

		assertEqual(t, d.Find(1), 1)
		assertEqual(t, d.Find(2), 2)
		assertEqual(t, d.Find(3), 2)
		assertEqual(t, d.Find(4), 2)
	})
}
