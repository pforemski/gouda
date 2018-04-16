/*
 * gouda/kmeans: a k-means implementation in Golang
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package kmeans

import "testing"
import "math/rand"
import "github.com/pforemski/gouda/point"

func TestSearch(t *testing.T) {
	points := make(point.Points, 0)

	for j := 0.0; j < 100.0; j += 10.0 {
		for i := 0; i < 10000; i++ {
			p := point.New(j+10.0*rand.Float64(), j+10.0*rand.Float64())
			points = append(points, p)
		}
	}

	clusters := Search(points, 10, 1000, 0.01)
	for i := range clusters {
		mean := clusters[i].Mean()
		sd := clusters[i].Stddev(mean)
		min := clusters[i].Min()
		max := clusters[i].Max()

		t.Logf("cluster[%d]:\n", i)
		t.Logf("  mean=%s, stddev=%s\n", mean, sd)
		t.Logf("  min=%s, max=%s\n", min, max)
	}
}
