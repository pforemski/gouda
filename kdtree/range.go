/*
 * gouda/kdtree/range: code implementing range search
 *
 * Copyright (C) 2018 Pawel Foremski <pjf@foremski.pl>
 * Licensed to you under GNU GPL v3
 */

package kdtree

import "github.com/pforemski/gopjf/point"

// search() returns all node's points that match given query range
// world is an information on what the node and all children can contain
func (node *KDNode) search(query *point.Range, world *point.Range, out point.Points) point.Points {
	// check if query intersects with world
	switch query.Intersects(world) {
	case 0: // no intersection
		return out

	case 1: // some intersection
		// check if node's point is contained
		if query.Contains(node.point) {
			out = append(out, node.point)
		}

		// check in the left child
		if node.left != nil {
			out = node.left.search(query, node.world_left(world), out)
		}

		// check in the right child
		if node.right != nil {
			out = node.right.search(query, node.world_right(world), out)
		}

	case 2: // fully contained
		out = node.dump(out)
	}

	return out
}

// world_left() updates worldview to reflect the left child of given node
func (parent *KDNode) world_left(world *point.Range) *point.Range {
	r := point.Range{}
	r.Min = world.Min // bottom limits don't change

	// upper limits: leave all but parent.axis
	r.Max = make([]float64, len(world.Max))
	copy(r.Max, world.Max)
	r.Max[parent.axis] = parent.point.V[parent.axis] // all elements < median

	return &r
}

// world_right() updates worldview to reflect the right child of given node
func (parent *KDNode) world_right(world *point.Range) *point.Range {
	r := point.Range{}
	r.Max = world.Max // upper limits don't change

	// botton limits: leave all but parent.axis
	r.Min = make([]float64, len(world.Min))
	copy(r.Min, world.Min)
	r.Min[parent.axis] = parent.point.V[parent.axis] // all elements >= median

	return &r
}
