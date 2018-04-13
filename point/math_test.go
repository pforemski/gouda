/*
 * gouda/kmeans: a k-means implementation in Golang
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "testing"

func TestOperations(t *testing.T) {
	p1 := New(0.25, 2.5)
	t.Logf("p1 = %s\n", p1)
	p1.Add(p1)
	t.Logf("p1 = p1+p1 = %s\n", p1)

	p2 := New(1.5, 5.0)
	t.Logf("p2 = %s\n", p2)
	p2.Add(p1)
	t.Logf("p2 = p1+p2 = %s\n", p2)

	p2.Mul(-2)
	t.Logf("p2 = -2 * p2 = %s\n", p2)
	t.Logf("p2.Min() = %g\n", p2.Min())
	t.Logf("p2.Max() = %g\n", p2.Max())

	points := Points{ p1, p2, New(10, 20) }
	t.Logf("points = %s\n", points)
	t.Logf("Min = %s\n", points.Min())
	t.Logf("Q1 = %s\n", points.Percentile(0.25))
	t.Logf("Mean() = %s\n", points.Mean())
	t.Logf("Q2 = Median() = %s\n", points.Median())
	t.Logf("Q3 = %s\n", points.Percentile(0.75))
	t.Logf("Max = %s\n", points.Max())
}
