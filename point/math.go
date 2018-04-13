/*
 * gouda/point: an n-dimensional point & tools
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "math"

// Euclidean() returns Euclidean distance between a and b; -1 on error
func (a *Point) Euclidean(b *Point) float64 {
	if len(a.V) != len(b.V) { return -1 }

	dist := 0.0
	for axis := 0; axis < len(a.V); axis++ {
		dist += math.Pow(a.V[axis] - b.V[axis], 2.0)
	}

	return math.Sqrt(dist)
}
// Add() adds p2 to p1
func (p1 *Point) Add(p2 *Point) {
	for i := 0; i < len(p1.V) && i < len(p2.V); i++ {
		p1.V[i] += p2.V[i]
	}
}

// Mul() multiplies all coordinates in p by fact
func (p *Point) Mul(fact float64) {
	for i := 0; i < len(p.V); i++ {
		p.V[i] *= fact
	}
}

// Mean() computes the arithmetic mean along each axis in points
func (points Points) Mean() *Point {
	axes := points[0].Axes()
	p := NewZero(axes)
	for i := range points {
		for axis := 0; axis < axes; axis++ {
			p.V[axis] += points[i].V[axis]
		}
	}

	fact := 1.0 / float64(len(points))
	for axis := 0; axis < axes; axis++ {
		p.V[axis] *= fact
	}

	return p
}

// Stddev() computes the standard deviation of points vs. given point
func (points Points) Stddev(mean *Point) *Point {
	axes := points[0].Axes()
	p := NewZero(axes)
	for i := range points {
		for axis := 0; axis < axes; axis++ {
			p.V[axis] += math.Pow(points[i].V[axis] - mean.V[axis], 2)
		}
	}

	fact := 1.0 / float64(len(points))
	for axis := 0; axis < axes; axis++ {
		p.V[axis] = math.Sqrt(p.V[axis] * fact)
	}

	return p
}

// Min() finds minimum value on each axis
func (points Points) Min() *Point {
	axes := points[0].Axes()

	// use first point as start
	p := Copy(points[0])
	p.D = nil

	for i := range points {
		for axis := 0; axis < axes; axis++ {
			if points[i].V[axis] < p.V[axis] {
				p.V[axis] = points[i].V[axis]
			}
		}
	}

	return p
}

// Max() finds maximum value on each axis
func (points Points) Max() *Point {
	axes := points[0].Axes()

	// use first point as start
	p := Copy(points[0])
	p.D = nil

	for i := range points {
		for axis := 0; axis < axes; axis++ {
			if points[i].V[axis] > p.V[axis] {
				p.V[axis] = points[i].V[axis]
			}
		}
	}

	return p
}
