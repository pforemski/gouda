/*
 * gouda/kmeans: a k-means implementation in Golang
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "testing"

func TestDistance(t *testing.T) {
	p1 := New(0.25, 2.5)
	p2 := New(1.5, 5.0)

	t.Logf("Euclidean(p1, p2) = %g\n", Euclidean(p1, p2))
	t.Logf("Maxdiff(p1, p2) = %g\n", Maxdiff(p1, p2))
	t.Logf("Taxicab(p1, p2) = %g\n", Taxicab(p1, p2))
}
