// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CAUTION: If this file is not vector.go, it was generated
// automatically from vector.go - DO NOT EDIT in that case!

package main

const (
	initialSize = 40
)

func (p *NodeVector) Less(i, j int) bool {
	a := *p
	return a[i].Less(&a[j])
}

func (p *NodeVector) realloc(length, capacity int) (b []Node) {
	if capacity < initialSize {
		capacity = initialSize
	}
	if capacity < length {
		capacity = length
	}
	b = make(NodeVector, length, capacity)
	copy(b, *p)
	*p = b
	return
}

// Insert n elements at position i.
func (p *NodeVector) Expand(i, n int) {
	a := *p

	// make sure we have enough space
	len0 := len(a)
	len1 := len0 + n
	if len1 <= cap(a) {
		// enough space - just expand
		a = a[0:len1]
	} else {
		// not enough space - double capacity
		capb := cap(a) * 2
		if capb < len1 {
			// still not enough - use required length
			capb = len1
		}
		// capb >= len1
		a = p.realloc(len1, capb)
	}

	// make a hole
	for j := len0 - 1; j >= i; j-- {
		a[j+n] = a[j]
	}

	*p = a
}

// Insert n elements at the end of a vector.
func (p *NodeVector) Extend(n int) { p.Expand(len(*p), n) }

// Resize changes the length and capacity of a vector.
// If the new length is shorter than the current length, Resize discards
// trailing elements. If the new length is longer than the current length,
// Resize adds the respective zero values for the additional elements. The capacity
// parameter is ignored unless the new length or capacity is longer than the current
// capacity. The resized vector's capacity may be larger than the requested capacity.
func (p *NodeVector) Resize(length, capacity int) *NodeVector {
	a := *p

	if length > cap(a) || capacity > cap(a) {
		// not enough space or larger capacity requested explicitly
		a = p.realloc(length, capacity)
	} else if length < len(a) {
		// clear trailing elements
		for i := range a[length:] {
			var zero Node
			a[length+i] = zero
		}
	}

	*p = a[0:length]
	return p
}

// Len returns the number of elements in the vector.
// Same as len(*p).
func (p *NodeVector) Len() int { return len(*p) }

// Cap returns the capacity of the vector; that is, the
// maximum length the vector can grow without resizing.
// Same as cap(*p).
func (p *NodeVector) Cap() int { return cap(*p) }

// At returns the i'th element of the vector.
func (p *NodeVector) At(i int) Node { return (*p)[i] }

// Set sets the i'th element of the vector to value x.
func (p *NodeVector) Set(i int, x Node) { (*p)[i] = x }

// Last returns the element in the vector of highest index.
func (p *NodeVector) Last() Node { return (*p)[len(*p)-1] }

// Copy makes a copy of the vector and returns it.
func (p *NodeVector) Copy() NodeVector {
	arr := make(NodeVector, len(*p))
	copy(arr, *p)
	return arr
}

// Insert inserts into the vector an element of value x before
// the current element at index i.
func (p *NodeVector) Insert(i int, x Node) {
	p.Expand(i, 1)
	(*p)[i] = x
}

// Delete deletes the i'th element of the vector.  The gap is closed so the old
// element at index i+1 has index i afterwards.
func (p *NodeVector) Delete(i int) {
	a := *p
	n := len(a)

	copy(a[i:n-1], a[i+1:n])
	var zero Node
	a[n-1] = zero // support GC, zero out entry
	*p = a[0 : n-1]
}

// InsertVector inserts into the vector the contents of the vector
// x such that the 0th element of x appears at index i after insertion.
func (p *NodeVector) InsertVector(i int, x *NodeVector) {
	b := *x

	p.Expand(i, len(b))
	copy((*p)[i:i+len(b)], b)
}

// Cut deletes elements i through j-1, inclusive.
func (p *NodeVector) Cut(i, j int) {
	a := *p
	n := len(a)
	m := n - (j - i)

	copy(a[i:m], a[j:n])
	for k := m; k < n; k++ { //TODO(bflm) don't zero out the elements unless it's a Vector.
		var zero Node
		a[k] = zero // support GC, zero out entries
	}

	*p = a[0:m]
}

// Slice returns a new sub-vector by slicing the old one to extract slice [i:j].
// The elements are copied. The original vector is unchanged.
func (p *NodeVector) Slice(i, j int) *NodeVector {
	var s NodeVector
	s.realloc(j-i, 0) // will fail in Init() if j < i
	copy(s, (*p)[i:j])
	return &s
}

// Convenience wrappers

// Push appends x to the end of the vector.
func (p *NodeVector) Push(x Node) { p.Insert(len(*p), x) }

// Pop deletes the last element of the vector.
func (p *NodeVector) Pop() Node {
	a := *p

	i := len(a) - 1
	x := a[i]
	var zero Node
	a[i] = zero // support GC, zero out entry
	*p = a[0:i]
	return x
}

// AppendVector appends the entire vector x to the end of this vector.
func (p *NodeVector) AppendVector(x *NodeVector) { p.InsertVector(len(*p), x) }

// Swap exchanges the elements at indexes i and j.
func (p *NodeVector) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
}

// Do calls function f for each element of the vector, in order.
// The behavior of Do is undefined if f changes *p.
func (p *NodeVector) Do(f func(elem Node)) {
	for _, e := range *p {
		f(e)
	}
}
