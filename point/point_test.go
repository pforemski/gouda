/*
 * gouda/kmeans: a k-means implementation in Golang
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "testing"

func TestString(t *testing.T) {
	p1 := New(0.15, 2.25)
	t.Logf("p1: %s\n", p1)

	p2 := NewD(p1, 0.753, -2.25)
	t.Logf("p2: %s\n", p2)
}
