# dsu 

[![GoDoc](https://godoc.org/github.com/ihebu/dsu?status.svg)](https://godoc.org/github.com/ihebu/dsu)
[![Go Report Card](https://goreportcard.com/badge/github.com/ihebu/dsu)](https://goreportcard.com/report/github.com/ihebu/dsu)
[![Build Status](https://travis-ci.com/ihebu/dsu.svg?branch=main)](https://travis-ci.com/ihebu/dsu)
[![codecov](https://codecov.io/gh/ihebu/dsu/branch/main/graph/badge.svg)](https://codecov.io/gh/ihebu/dsu)

Implementation of the Disjoint-Set data structure.
The Disjoint-Set, Also called a Union-Find or Merge-Find set, is a data structure that stores a collection of disjoint (non-overlapping) sets. Equivalently, it stores a partition of a set into disjoint subsets. It provides operations for adding new sets,
merging sets (replacing them by their union), and finding a representative member of a set. The last operation allows to find out efficiently if any two elements are in the same or different sets.

## Installation 

```bash
go get github.com/ihebu/dsu
```

## Documentation 

You can check the code documentation [here](https://godoc.org/github.com/ihebu/dsu)

## Usage Example

```go

// Create a new disjoint-set
d := dsu.New()

// Add the elements 1, 2, 3 to the set
d.Add(1)
d.Add(2)
d.Add(3)

// The set is now {1}, {2}, {3}

// Unite the sets {1}, {2} 
d.Union(1, 2)

// The set is now {1, 2}, {3}

// Find the representative element of each set
d.Find(1) // returns 2
d.Find(2) // returns 2
d.Find(3) // returns 3

// Check the existence of an element in the set
d.Contains(2) // returns true
d.Contains(54) // returns false

// Note : you can add elements of different type in the set
// Example

d.Add("hello")
d.Add(34.5)

d.Union("hello", 34.5)
```
