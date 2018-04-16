/*
 * gouda/point: an n-dimensional point & tools
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "math"

// Euclidean() returns Euclidean distance between a and b
func Euclidean(a, b *Point) float64 {
	dist := 0.0
	for axis := 0; axis < len(a.V); axis++ {
		dist += math.Pow(a.V[axis] - b.V[axis], 2.0)
	}
	return math.Sqrt(dist)
}

// Maxdiff() returns the maximum difference between a and b on any axis
func Maxdiff(a, b *Point) float64 {
	dist := 0.0
	for axis := 0; axis < len(a.V); axis++ {
		diff := math.Abs(a.V[axis] - b.V[axis])
		if diff > dist { dist = diff }
	}
	return dist
}

// Taxicab() returns the taxicab distance between a and b;
// see https://en.wikipedia.org/wiki/Taxicab_geometry
func Taxicab(a, b *Point) float64 {
	dist := 0.0
	for axis := 0; axis < len(a.V); axis++ {
		dist += math.Abs(a.V[axis] - b.V[axis])
	}
	return dist
}
