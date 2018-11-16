/*
 * gouda/interpolation/lagrange: Lagrange interpolation polynomials for 2D points
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package interpolate

import (
	"fmt"
	"github.com/pforemski/gouda/point"
)

// Lagrange represents internal state of Lagrange interpolation polynomial
type Lagrange struct {
	// interpolation points (must be 2D)
	Points      point.Points

	// denominator values of the n+1 auxilliary polynomials
	wd          []float64
}

// New() returns new Lagrange interpolation for given set of 2D points
func New(points point.Points) (*Lagrange, error) {
	lg := &Lagrange{}
	lg.Points = points

	// check points, pre-compute parts of aux. polynomials
	lg.wd = make([]float64, len(points))
	for i := range lg.Points {
		if len(lg.Points[i].V) != 2 {
			return nil, fmt.Errorf("point #%d (%s): must be 2D", i, lg.Points[i])
		}

		lg.wd[i] = 1.0
		for j := range lg.Points {
			if i != j { lg.wd[i] *= lg.Points[i].V[0] - lg.Points[j].V[0] }
		}

		if lg.wd[i] == 0 {
			return nil, fmt.Errorf("point #%d (%s): repetition (wd=0)", i, lg.Points[i])
		}
	}

	return lg, nil
}

// w() returns the value of i-th aux polynomial at point x
func (lg *Lagrange) w(i int, x float64) float64 {
	v := 1.0
	for j := range lg.Points {
		if i != j { v *= x - lg.Points[j].V[0] }
	}
	return v / lg.wd[i]
}

// Interpolate() predicts function value in point x
func (lg *Lagrange) Interpolate(x float64) float64 {
	v := 0.0
	for i := range lg.Points {
		v += lg.w(i, x) * lg.Points[i].V[1]
	}
	return v
}