package dsu

import (
	"math/rand"
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

	d.Add(1)
	d.Add("foo")

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

func TestUnion(t *testing.T) {
	t.Run("Union of non existing elements", func(t *testing.T) {
		d := New()
		d.Add(1)

		assertEqual(t, d.Union(1, 2), false)
		assertEqual(t, d.Union(2, 3), false)
	})

	t.Run("Union of existing elements", func(t *testing.T) {
		d := New()
		node1 := &node{value: 1, parent: nil, size: 1}
		node2 := &node{value: 2, parent: nil, size: 2}
		node3 := &node{value: 3, parent: node2, size: 1}
		node4 := &node{value: 4, parent: nil, size: 1}

		d.nodes[1] = node1
		d.nodes[2] = node2
		d.nodes[3] = node3
		d.nodes[4] = node4

		assertEqual(t, d.Union(2, 3), false)
		assertEqual(t, d.Union(1, 3), true)
		assertEqual(t, d.Union(1, 3), false)
		assertEqual(t, d.Union(3, 4), true)

		assertEqual(t, d.nodes[1].size, 1)
		assertEqual(t, d.nodes[3].size, 1)
		assertEqual(t, d.nodes[2].size, 4)
		assertEqual(t, d.nodes[4].size, 1)
	})
}

var result bool

func BenchmarkContains(b *testing.B) {
	rand.Seed(42)
	d := New()
	for i := 0; i < 100000; i++ {
		d.Add(rand.Intn(100000))
	}
	var r bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = d.Contains(rand.Intn(100000))
	}
	result = r
}

func BenchmarkAdd(b *testing.B) {
	rand.Seed(42)
	d := New()
	var r bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = d.Add(rand.Intn(100000))
	}
	result = r
}

func BenchmarkUnion(b *testing.B) {
	rand.Seed(42)
	d := New()
	for i := 0; i < 100000; i++ {
		d.Add(i)
	}
	var r bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := rand.Intn(100000)
		y := rand.Intn(100000)
		r = d.Union(x, y)
	}
	result = r
}

var result2 interface{}

func BenchmarkFind(b *testing.B) {
	rand.Seed(42)
	d := New()
	for i := 0; i < 100000; i++ {
		d.Add(i)
	}
	for i := 0; i < 1000; i++ {
		x := rand.Intn(100000)
		y := rand.Intn(100000)
		d.Union(x, y)
	}
	var r interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = d.Find(rand.Intn(100000))
	}
	result2 = r
}
