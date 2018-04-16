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
		for i := 0; i < 1000; i++ {
			p := point.New(j+10.0*rand.Float64(), j+10.0*rand.Float64())
			points = append(points, p)
		}
	}

	dist := []func(*point.Point,*point.Point)float64{
		point.Euclidean,
		point.Maxdiff,
		point.Taxicab,
	}

	for di := range dist {
		t.Logf("dist fn %d\n", di)

		clusters := SearchDist(points, 10, 100, 0.01, dist[di])
		for i := range clusters {
			mean := clusters[i].Mean()
			sd := clusters[i].Stddev(mean)
			min := clusters[i].Min()
			max := clusters[i].Max()

			t.Logf("cluster[%d]:\n", i)
			t.Logf("  mean=%s, stddev=%s\n", mean, sd)
			t.Logf("  min=%s, max=%s\n", min, max)
		}

		t.Logf("\n")
	}
}
