/*
 * gouda/point: an n-dimensional point & tools
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "math"

// Range represents a range in n-dimensional space (a rectangle)
type Range struct {
	Min []float64
	Max []float64
}

// NewInfiniteRange() creates a range that contains everything
// Parameter axes specifies dimensionality
func NewInfiniteRange(axes int) *Range {
	r := Range{}
	r.Min = make([]float64, axes)
	r.Max = make([]float64, axes)

	for axis := 0; axis < axes; axis++ {
		r.Min[axis] = math.Inf(-1)
		r.Max[axis] = math.Inf(1)
	}

	return &r
}

// Contains() returns true if given point is within the query
func (query *Range) Contains(point *Point) bool {
	for axis := 0; axis < len(point.V); axis++ {
		if point.V[axis] < query.Min[axis] { return false }
		if point.V[axis] > query.Max[axis] { return false }
	}
	return true
}

// Intersects() checks intersection between query and given world range
// 0 means no intersection, 1 means partial, 2 means query fully contains r
func (query *Range) Intersects(world *Range) int {
	ret := 2

	for axis := 0; axis < len(query.Min); axis++ {
		if query.Min[axis] >= world.Max[axis] { return 0 } // no intersection possible
		if query.Max[axis] <  world.Min[axis] { return 0 } // no intersection possible

		if ret == 2 {
			if query.Max[axis] < world.Max[axis] { ret = 1 } // partial intersection
			if query.Min[axis] > world.Min[axis] { ret = 1 } // partial intersection
		}
	}

	return ret
}


