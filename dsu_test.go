package dsu

import "testing"

func assertEqual(t *testing.T, got, expected bool) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %t, got %t", expected, got)
	}
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
