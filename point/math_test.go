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

	points := Points{ p1, p2 }
	t.Logf("Mean(p1,p2) = %s\n", points.Mean())
}
