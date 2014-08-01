package gogl

import (
	"gopkg.in/fatih/set.v0"
)

// Shared helper function for edge lists to enumerate vertices.
func esEachVertex(el interface{}, fn VertexLambda) {
	set := set.NewNonTS()

	el.(EdgeEnumerator).EachEdge(func(e Edge) (terminate bool) {
		set.Add(e.Both())
		return
	})

	for _, v := range set.List() {
		if fn(v) {
			return
		}
	}
}

// Shared helper function for edge lists to report vertex count.
func esOrder(el interface{}) int {
	set := set.NewNonTS()

	el.(EdgeEnumerator).EachEdge(func(e Edge) (terminate bool) {
		set.Add(e.Both())
		return
	})

	return set.Size()
}

// An EdgeList is a naive GraphSource implementation that is backed only by an edge slice.
//
// EdgeLists are primarily intended for use as fixtures.
//
// It is inherently impossible for an EdgeList to represent a vertex isolate (degree 0) fully
// correctly, as vertices are only described in the context of their edges. One can be a
// little hacky, though, and represent one with a loop. As gogl expects graph implementations
// to simply discard loops if they are disallowed by the graph's constraints (i.e., in simple
// and multigraphs), they *should* be interpreted as vertex isolates.
type EdgeList []Edge

func (el EdgeList) EachVertex(fn VertexLambda) {
	esEachVertex(el, fn)
}

func (el EdgeList) Order() int {
	return esOrder(el)
}

func (el EdgeList) EachEdge(fn EdgeLambda) {
	for _, e := range el {
		if fn(e) {
			return
		}
	}
}

// A WeightedEdgeList is a naive GraphSource implementation that is backed only by an edge slice.
//
// This variant is for weighted edges.
type WeightedEdgeList []WeightedEdge

func (el WeightedEdgeList) EachVertex(fn VertexLambda) {
	esEachVertex(el, fn)
}

func (el WeightedEdgeList) Order() int {
	return esOrder(el)
}

func (el WeightedEdgeList) EachEdge(fn EdgeLambda) {
	for _, e := range el {
		if fn(e) {
			return
		}
	}
}

// A LabeledEdgeList is a naive GraphSource implementation that is backed only by an edge slice.
//
// This variant is for labeled edges.
type LabeledEdgeList []LabeledEdge

func (el LabeledEdgeList) EachVertex(fn VertexLambda) {
	esEachVertex(el, fn)
}

func (el LabeledEdgeList) Order() int {
	return esOrder(el)
}

func (el LabeledEdgeList) EachEdge(fn EdgeLambda) {
	for _, e := range el {
		if fn(e) {
			return
		}
	}
}

// A DataEdgeList is a naive GraphSource implementation that is backed only by an edge slice.
//
// This variant is for labeled edges.
type DataEdgeList []DataEdge

func (el DataEdgeList) EachVertex(fn VertexLambda) {
	esEachVertex(el, fn)
}

func (el DataEdgeList) Order() int {
	return esOrder(el)
}

func (el DataEdgeList) EachEdge(fn EdgeLambda) {
	for _, e := range el {
		if fn(e) {
			return
		}
	}
}