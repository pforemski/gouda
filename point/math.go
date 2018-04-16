/*
 * gouda/point: an n-dimensional point & tools
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "math"
import "sort"

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
	if len(points) == 0 { return NewZero(0) }

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

// Percentile() finds given percentile on each axis
func (points Points) Percentile(percentile float64) *Point {
	if len(points) == 0 { return NewZero(0) }

	axes := points[0].Axes()
	p := NewZero(axes)

	for axis := 0; axis < axes; axis++ {
		vals := make([]float64, len(points))
		for i := range points {
			vals[i] = points[i].V[axis]
		}
		sort.Float64s(vals)

		// use precise location
		l := float64(len(vals) - 1) // last val
		i, f := math.Modf(l * percentile)
		if f < 0.01 || i == l {
			p.V[axis] = vals[int(i)]
		} else {
			p.V[axis] = (1.0-f)*vals[int(i)] + f*vals[int(i)+1]
		}
	}

	return p
}

// Median() finds median using Percentile(0.5)
func (points Points) Median() *Point { return points.Percentile(0.5); }

// Stddev() computes the standard deviation of points vs. given point
func (points Points) Stddev(mean *Point) *Point {
	if len(points) == 0 { return NewZero(0) }

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

// Min() finds the minimum value on each axis; it is not "the minimal point" of all points
func (points Points) Min() *Point {
	if len(points) == 0 { return NewZero(0) }

	// use first point as start
	p := points[0].Copy()
	p.D = nil

	for i := 1; i < len(points); i++ {
		for axis := 0; axis < len(p.V); axis++ {
			if points[i].V[axis] < p.V[axis] {
				p.V[axis] = points[i].V[axis]
			}
		}
	}

	return p
}

// Min() finds the minimum value on any axis
func (p *Point) Min() float64 {
	min := p.V[0]
	for i := 1; i < len(p.V); i++ {
		if p.V[i] < min { min = p.V[i] }
	}
	return min
}

// Max() finds the maximum value on each axis; it is not "the maximum point" of all points
func (points Points) Max() *Point {
	if len(points) == 0 { return NewZero(0) }

	// use first point as start
	p := points[0].Copy()
	p.D = nil

	for i := 1; i < len(points); i++ {
		for axis := 0; axis < len(p.V); axis++ {
			if points[i].V[axis] > p.V[axis] {
				p.V[axis] = points[i].V[axis]
			}
		}
	}

	return p
}

// Max() finds the maximum value on any axis
func (p *Point) Max() float64 {
	max := p.V[0]
	for i := 1; i < len(p.V); i++ {
		if p.V[i] > max { max = p.V[i] }
	}
	return max
}
