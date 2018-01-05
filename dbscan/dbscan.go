/*
 * gouda/dbscan: a DBSCAN implementation in Golang
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package dbscan

import "github.com/pforemski/gopjf/point"
import "github.com/pforemski/gopjf/kdtree"

// Search() runs DBSCAN algorithm on given points
//
// eps is the maximum distance to neighbors, in all dimensions
// minpoints is the minimum number of points in eps neighborhood to form a core point
// clusters is a slice of Points, where each element corresponds to a cluster
// clusters[0] holds noise, clusters[1] holds cluster #1, etc.
func Search(points point.Points, eps []float64, minpoints int) (clusters []point.Points) {
	// needed for algo
	C := 0
	labels := make(map[*point.Point]int)  // -1 means noise, 0 undefined, >0 cluster label
	kdt := kdtree.New(points)

	// try all points
	for _, P := range points {
		// already visited?
		if labels[P] != 0 { continue }

		// is it a core point?
		neighbors := kdt.Search(P, eps)
		if len(neighbors) < minpoints { // no
			labels[P] = -1
			continue
		}

		// yes, we found a cluster
		C++
		labels[P] = C

		// for de-duplication
		dedup := make(map[*point.Point]bool)
		for _, N := range neighbors { dedup[N] = true }

		// go through each neighbor, try expanding
		for i := 0; i < len(neighbors); i++ {
			N := neighbors[i]

			switch labels[N] {
			case -1: // noise? change to border point
				labels[N] = C
				continue

			case 0:  // not already visited? try expanding
				labels[N] = C

				// is this neighbor a core point?
				neighbors2 := kdt.Search(N, eps)
				if len(neighbors2) < minpoints { // no
					continue
				}

				// try expanding
				for _, N := range neighbors2 {
					if dedup[N] { continue } // already tried for this cluster
					neighbors = append(neighbors, N)
				}

			default: // already in a cluster, skip
				continue
			}
		}
	}

	// prepare return value
	clusters = make([]point.Points, C+1)
	for label := 0; label <= C; label++ {
		clusters[label] = make(point.Points, 0)
	}

	// rewrite from labels to clusters
	for point, label := range labels {
		if label == 0 { panic("unvisited point") }
		if label == -1 { label = 0 }

		clusters[label] = append(clusters[label], point)
	}

	return
}
