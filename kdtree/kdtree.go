/*
 * gouda/kdtree: a kd-tree implementation in Golang
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package kdtree

import "github.com/pforemski/gouda/point"

// KDNode represents a kd-tree
type KDNode struct {
	point   *point.Point // point stored at node
	axis    int          // splitting axis
	left    *KDNode      // left child:  points less than this point (on axis)
	right   *KDNode      // right child: points greater than this point (on axis)
	count   int          // number of all points stored in node and both children
}

// New() creates a new kd-tree with given points
func New(points point.Points) *KDNode {
	return insert(points, 0)
}

// Search() performs range search for [ref - margin, ref + margin],
// starting at given kd-tree node, returning a slice of pointers to matching points
func (node *KDNode) Search(ref *point.Point, margin []float64) point.Points {
	// translate ref+margin into range
	query := point.NewInfiniteRange(len(ref.V))
	for axis := 0; axis < len(ref.V) && axis < len(margin); axis++ {
		if margin[axis] >= 0 {
			query.Min[axis] = ref.V[axis] - margin[axis]
			query.Max[axis] = ref.V[axis] + margin[axis]
		}
	}

	// prepare info on current's node worldview
	world := point.NewInfiniteRange(len(ref.V))

	// query
	points := make(point.Points, 0, 32) // NB: pre-allocate for 32 results
	points = node.search(query, world, points)

	return points
}

// Dump() returns a slice of pointers to all points stored in a given kd-tree node,
// and all of it's children (left / right)
func (node *KDNode) Dump() point.Points {
	points := make(point.Points, 0, node.count)
	return node.dump(points)
}

// ------------------------------------------

func (node *KDNode) dump(points point.Points) point.Points {
	points = append(points, node.point)
	if node.left  != nil { points = node.left.dump(points) }
	if node.right != nil { points = node.right.dump(points) }
	return points
}

func insert(points point.Points, depth int) *KDNode {
	if len(points) == 0 { return nil }

	node := &KDNode{
		axis: (depth % len(points[0].V)),
		count: len(points),
	}

	// find median by sampling on given axis
	median := points.SampleMedian(node.axis, 250)

	// divide
	points_below := make(point.Points, 0, len(points)/2)
	points_above := make(point.Points, 0, len(points)/2)
	for i := range points {
		if points[i] == median {
			continue
		} else if points[i].V[node.axis] < median.V[node.axis] {
			points_below = append(points_below, points[i])
		} else {
			points_above = append(points_above, points[i])
		}
	}

	node.point = median
	node.left = insert(points_below, depth + 1)
	node.right = insert(points_above, depth + 1)

	return node
}
