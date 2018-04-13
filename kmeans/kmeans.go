/*
 * gouda/kmeans: a k-means implementation in Golang
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package kmeans

import "github.com/pforemski/gouda/point"

// Search() runs k-means algorithm on given points
//
// k is the number of clusters;
// max_iter is maximum number of iterations;
// min_change is minimum change in centers (any axis) to continue;
//
// clusters is a slice of Points, where each element corresponds to a cluster
func Search(points point.Points, k int, max_iter int, min_change float64) (clusters []point.Points) {
	// prepare return value
	clusters = make([]point.Points, k)
	for label := 0; label < k; label++ {
		clusters[label] = make(point.Points, 0)
	}

	// sanity checks
	if len(points) < 1 { return }
	axes := len(points[0].V)

	// begin with simple state
	assignments := make([]int, len(points))
	centers := points.Sample(k)

	// rock and roll
	for iteration := 0; iteration < max_iter; iteration++ {
		// re-assign to closest clusters
		for i := range points {
			min_dist := -1.0
			for c := range centers {
				dist := points[i].Euclidean(centers[c])
				if min_dist < 0 || dist < min_dist {
					assignments[i] = c
					min_dist = dist
				}
			}
		}

		// re-compute centers, check difference vs. old centers
		new_centers := make(point.Points, k)
		new_counter := make([]float64, k)
		max_change := 0.0
		for c := range new_centers {
			new_centers[c] = point.NewZero(axes)
		}
		for i := range assignments {
			c := assignments[i]
			new_centers[c].Add(points[i])
			new_counter[c] += 1
		}
		for c := range new_centers {
			new_centers[c].Mul(1.0/new_counter[c])
			diff := centers[c].Euclidean(new_centers[c])
			if diff > max_change { max_change = diff }
		}

		// how much did they change?
		if max_change < min_change {
			break
		}

		centers = new_centers
	}

	// re-write the results
	for i := range assignments {
		c := assignments[i]
		clusters[c] = append(clusters[c], points[i])
	}

	return
}
