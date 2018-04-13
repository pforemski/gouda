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

// Points is a collection of memory pointers to points
type Points []*Point

// New() creates a new point
func New(vals ...float64) *Point {
	ret := &Point{}
	ret.V = vals
	return ret
}

// NewD() creates a new point with data D
func NewD(D interface{}, vals ...float64) *Point {
	ret := &Point{}
	ret.V = vals
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
func Copy(p *Point) *Point {
	ret := &Point{}
	ret.V = make([]float64, len(p.V))
	copy(ret.V, p.V)
	ret.D = p.D
	return ret
}

// Axes() returns number of axes in p
func (p *Point) Axes() int { return len(p.V) }

// String() converts a Point to a string object
func (p *Point) String() (ret string) {
	ret = fmt.Sprintf("%.4g", p.V)
	if p.D != nil { ret += fmt.Sprintf("->(%T %v)", p.D, p.D) }
	return
}
