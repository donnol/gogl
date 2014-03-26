package gogl

/* Vertex structures */

// As a rule, gogl tries to place as low a requirement on its vertices as
// possible. This is because, from a purely graph theoretic perspective,
// vertices are inert. Boring, even. Graphs are more about the topology, the
// characteristics of the edges connecting the points than the points
// themselves. Your use case cares about the content of your vertices, but gogl
// does not.  Consequently, anything can act as a vertex.
type Vertex interface{}

/* Edge structures */

// A graph's behaviors is primarily a product of the constraints and
// capabilities it places on its edges. These constraints and capabilities
// determine whether certain types of operations are possible on the graph, as
// well as the efficiencies for various operations.

// gogl aims to provide a diverse range of graph implementations that can meet
// the varying constraints and implementation needs, but still achieve optimal
// performance given those constraints.

type Edge interface {
	Source() Vertex
	Target() Vertex
	Both() (Vertex, Vertex)
}

type WeightedEdge interface {
	Edge
	Weight() int
}

// BaseEdge is a struct used internally to represent edges and meet the Edge
// interface requirements. It uses the standard notation, (u,v), for vertex
// pairs in an edge.
type BaseEdge struct {
	U Vertex
	V Vertex
}

func (e BaseEdge) Source() Vertex {
	return e.U
}

func (e BaseEdge) Target() Vertex {
	return e.V
}

func (e BaseEdge) Both() (Vertex, Vertex) {
	return e.U, e.V
}

type BaseWeightedEdge struct {
	BaseEdge
	W int
}

func (e BaseWeightedEdge) Weight() int {
	return e.W
}

/* Graph structures */

// Graph is gogl's most basic interface: it contains only the methods that
// *every* type of graph implements.
//
// Graph is intentionally underspecified: both directed and undirected graphs
// implement it; simple graphs, multigraphs, weighted, labeled, or any
// combination thereof.
//
// The semantics of some of these methods vary slightly from one graph type
// to another, but in general, the baisc Graph methods are supplemented, not
// superceded, by the methods in more specific interfaces.
//
// Graph is a purely read oriented interface; the various Mutable*Graph
// interfaces contain the methods for writing.
type Graph interface {
	EachVertex(f func(vertex Vertex))
	EachEdge(f func(edge Edge))
	EachAdjacent(vertex Vertex, f func(adjacent Vertex))
	HasVertex(vertex Vertex) bool
	HasEdge(e Edge) bool
	Order() int
	Size() int
	InDegree(vertex Vertex) (int, bool)
	OutDegree(vertex Vertex) (int, bool)
}

// DirectedGraph describes a Graph all of whose edges are directed.
//
// Implementing DirectedGraph is the only unambiguous signal gogl provides
// that a graph's edges are directed.
type DirectedGraph interface {
	Graph
	Transpose() DirectedGraph
}

// MutableGraph describes a graph with basic edges (no weighting, labeling, etc.)
// that can be modified freely by adding or removing vertices or edges.
type MutableGraph interface {
	Graph
	EnsureVertex(vertices ...Vertex)
	RemoveVertex(vertices ...Vertex)
	AddEdges(edges ...Edge)
	RemoveEdges(edges ...Edge)
}

// A simple graph is in opposition to a multigraph: it disallows loops and
// parallel edges.
type SimpleGraph interface {
	Graph
	Density() float64
}

// A weighted graph is a graph subtype where the edges have a numeric weight;
// as described by the WeightedEdge interface, this weight is a signed int.
//
// WeightedGraphs have both the HasEdge() and HasWeightedEdge() methods.
// Correct implementations should treat the difference as a matter of
// strictness:
//
// HasEdge() should return true as long as an edge exists
// connecting the two given vertices (respecting directed or undirected as
// appropriate), regardless of its weight.
//
// HasWeightedEdge() should return true iff an edge exists connecting the
// two given vertices (respecting directed or undirected as appropriate),
// AND if the edge weights are the same.
type WeightedGraph interface {
	Graph
	HasWeightedEdge(e WeightedEdge) bool
	EachWeightedEdge(f func(edge WeightedEdge))
}

// MutableWeightedGraph is the mutable version of a weighted graph. Its
// AddEdges() method is incompatible with MutableGraph, guaranteeing
// only weighted edges can be present in the graph.
type MutableWeightedGraph interface {
	WeightedGraph
	EnsureVertex(vertices ...Vertex)
	RemoveVertex(vertices ...Vertex)
	AddEdges(edges ...WeightedEdge)
	RemoveEdges(edges ...WeightedEdge)
}
