/*
 * gouda/kmeans: a k-means implementation in Golang
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package kmeans

import "github.com/pforemski/gouda/point"

// Search() runs SearchDist() using the point.Euclidean as distance function
//
// k is the number of clusters;
// max_iter is maximum number of iterations;
// min_change is minimum change in centers (any axis) to continue;
//
// clusters is a slice of Points, where each element corresponds to a cluster
func Search(points point.Points, k int, max_iter int, min_change float64) (clusters []point.Points) {
	return SearchDist(points, k, max_iter, min_change, point.Euclidean)
}

// SearchDist() runs the k-means algorithm using given distance function
//
// k is the number of clusters;
// max_iter is maximum number of iterations;
// min_change is minimum change in centers (any axis) to continue;
// dist_func is the distance function;
//
// clusters is a slice of Points, where each element corresponds to a cluster
func SearchDist(points point.Points, k int, max_iter int, min_change float64,
	dist_func func(*point.Point, *point.Point) float64) (clusters []point.Points) {
	// prepare return value
	clusters = make([]point.Points, k)
	for label := 0; label < k; label++ {
		clusters[label] = make(point.Points, 0)
	}

	// sanity checks
	if len(points) < 1 { return }

	// begin with simple state
	assignments := make([]int, len(points))
	centers := points.RandomSample(k).Copy()

	// prepare for future loops
	next_centers := point.NewZeros(k, centers[0].Axes())
	next_counter := make([]float64, k)

	// rock and roll
	for iteration := 0; iteration < max_iter || max_iter == 0; iteration++ {
		// clear state
		changes := 0
		next_centers.Zeros()
		for i := range next_counter { next_counter[i] = 0 }

		// assign each point to closest cluster
		for i := range points {
			min_dist := -1.0
			min_clus := -1

			// search for the closest cluster
			for c := range centers {
				dist := dist_func(points[i], centers[c])
				if min_clus < 0 || dist < min_dist {
					min_dist = dist
					min_clus = c
				}
			}

			// assignment changed?
			if min_clus != assignments[i] {
				changes += 1
			}

			// assign
			assignments[i] = min_clus

			// track upcoming new cluster centers
			next_centers[min_clus].Add(points[i])
			next_counter[min_clus] += 1
		}

		// no more changes? quit
		if changes == 0 {
			break
		}

		// re-compute centers
		max_change := 0.0
		for c := range centers {
			if next_counter[c] > 0 {
				next_centers[c].Mul(1.0/next_counter[c])
			}

			// check how much it moved?
			if min_change > 0 {
				diff := dist_func(centers[c], next_centers[c])
				if diff > max_change { max_change = diff }
			}
		}

		// the change was too small?
		if min_change > 0 && max_change < min_change {
			break
		}

		// exchange the centers
		centers, next_centers = next_centers, centers
	}

	// re-write the results
	for i := range assignments {
		c := assignments[i]
		clusters[c] = append(clusters[c], points[i])
	}

	return
}
