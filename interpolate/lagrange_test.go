/*
 * gouda/interpolation/lagrange: Lagrange interpolation polynomials for 2D points
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package interpolate

import (
	"math"
	"testing"

	"github.com/pforemski/gouda/point"
)

func TestLagrange_Interpolate(t *testing.T) {
	maxerr := 1e-6

	tests := []struct {
		points  point.Points
		args    []float64
		want    []float64
	}{
	{ // #0: linear
		point.Points{
			point.New(0.0,    0),
			point.New(1.25,  12.5),
			point.New(6.11,  61.1),
			point.New(7.5,   75),
			point.New(10.1, 101),
		},
		[]float64{ -10.1, 0.5, 3.1,  5,  8,  200 },
		[]float64{ -101,    5,  31, 50, 80, 2000 },
	},
	}

	for ti, tt := range tests {
		if len(tt.args) != len(tt.want) {
			t.Fatalf("%d: len(args)!=len(want)", ti)
		}

		lg, err := New(tt.points)
		if err != nil {
			t.Fatalf("%d: New(): %s", ti, err)
		}

		for i,x := range tt.args {
			y := lg.Interpolate(x)
			want := tt.want[i]
			if math.Abs(y - want) > maxerr {
				t.Errorf("%d/%d: Interpolate(%g): %g != %g", ti, i, x, y, want)
			}
		}
	}
}
