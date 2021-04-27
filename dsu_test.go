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
