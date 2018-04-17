/*
 * gouda/point: an n-dimensional point & tools
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "fmt"

// Point represents an abstract point in n-dimensional space
type Point struct {
	// point coordinates in all dimensions
	V []float64

	// auxilliary data
	D interface{}
}

// New() creates a new point
func New(vals ...float64) *Point {
	ret := &Point{}
	if vals != nil {
		ret.V = vals
	} else {
		ret.V = make([]float64, 0)
	}
	return ret
}

// NewD() creates a new point with data D
func NewD(D interface{}, vals ...float64) *Point {
	ret := &Point{}
	if vals != nil {
		ret.V = vals
	} else {
		ret.V = make([]float64, 0)
	}
	ret.D = D
	return ret
}

// NewZero() creates a new zero point with given number of axes
func NewZero(axes int) *Point {
	ret := &Point{}
	ret.V = make([]float64, axes)
	return ret
}

// Copy() creates a new point by copying another one
func (p *Point) Copy() *Point {
	ret := &Point{}
	ret.V = make([]float64, len(p.V))
	copy(ret.V, p.V)
	ret.D = p.D
	return ret
}

// Axes() returns the number of axes in p
func (p *Point) Axes() int { return len(p.V) }

// String() converts a Point to a string object
func (p *Point) String() (ret string) {
	ret = fmt.Sprintf("%.4g", p.V)
	if p.D != nil { ret += fmt.Sprintf("->(%T %v)", p.D, p.D) }
	return
}

/********************************************************/

// Points is a collection of memory pointers to points
type Points []*Point

// NewZeros() creates new n zero Points with given number of axes
func NewZeros(n, axes int) Points {
	ret := make(Points, n)
	for i := range ret {
		ret[i] = NewZero(axes)
	}
	return ret
}

// Zeros() zero-es given slice of points
func (points Points) Zeros() {
	for i := range points {
		points[i].Zero()
	}
}

// Copy() creates new slice of points by copying another one
func (points Points) Copy() Points {
	ret := make(Points, len(points))
	for i := range points {
		ret[i] = points[i].Copy()
	}
	return ret
}
