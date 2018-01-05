/*
 * gouda/point: an n-dimensional point & tools
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

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
func NewData(D interface{}, vals ...float64) *Point {
	ret := &Point{}
	ret.V = vals
	ret.D = D
	return ret
}
