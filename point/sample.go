/*
 * gouda/point: an n-dimensional point & tools
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package point

import "sort"
import "math/rand"
import "time"

type sample_points struct {
	points Points
	axis   int
}

func (a sample_points) Len() int {
	return len(a.points)
}

func (a sample_points) Swap(i,j int) {
	a.points[i],a.points[j] = a.points[j],a.points[i]
}

func (a sample_points) Less(i,j int) bool {
	return a.points[i].V[a.axis] < a.points[j].V[a.axis]
}

// SampleMedian() finds median on given axis using uniform sampling
func (points Points) SampleMedian(axis int, sample_size int) *Point {
	sample := sample_points{ axis: axis }

	// how big sample?
	size := len(points)
	if size < 3 { return points[size-1] }
	if size > sample_size { size = sample_size }
	sample.points = make(Points, 0, size)

	// how big step through points?
	plen := float64(len(points))
	step := plen / float64(size)

	// take sample
	for fi := 0.0; fi < plen; fi += step {
		sample.points = append(sample.points, points[int(fi)])
	}

	// sort it
	sort.Sort(sample)

	// take median
	return sample.points[size / 2]
}

// Sample() returns a uniform sample of points
func (points Points) Sample(sample_size int) Points {
	// sanity checks
	if sample_size >= len(points) { return points }
	if sample_size < 1 { return Points{} }

	// take sample
	sample := make(Points, 0, sample_size)
	fpop := float64(len(points))
	step := fpop / float64(sample_size)
	for fi := step / 2.0; fi < fpop; fi += step {
		sample = append(sample, points[int(fi)])
	}

	// border case
	if len(sample) > sample_size {
		sample = sample[:len(sample)-1]
	}

	return sample
}

// RandomSample() returns a random sample of points
func (points Points) RandomSample(sample_size int) Points {
	// sanity checks
	if sample_size >= len(points) { return points }
	if sample_size < 1 { return Points{} }

	// take sample
	sample := make(Points, 0, sample_size)
	for _,i := range rand.Perm(len(points))[:sample_size] {
		sample = append(sample, points[i])
	}

	return sample
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
