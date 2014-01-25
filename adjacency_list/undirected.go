package adjacency_list

import (
	. "github.com/sdboyer/gogl"
	"github.com/fatih/set"
)

type Undirected struct {
	adjacencyList
}

func NewUndirected() *Undirected {
	list := &Undirected{}
	// Cannot assign to promoted fields in a composite literals.
	list.list = make(map[Vertex]VertexSet)

	// Type assertions to ensure interfaces are met
	var _ Graph = list
	var _ SimpleGraph = list
	var _ MutableGraph = list

	return list
}

// Creates a new Undirected graph from an edge set.
func NewUndirectedFromEdgeSet(set []Edge) *Undirected {
	g := NewUndirected()

	for _, edge := range set {
		g.addEdge(edge)
	}

	return g
}

// Returns the outdegree of the provided vertex. If the vertex is not present in the
// graph, the second return value will be false.
func (g *Undirected) OutDegree(vertex Vertex) (degree int, exists bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if exists = g.hasVertex(vertex); exists {
		degree = len(g.list[vertex])
	}
	return
}

// Returns the indegree of the provided vertex. If the vertex is not present in the
// graph, the second return value will be false.
func (g *Undirected) InDegree(vertex Vertex) (degree int, exists bool) {
	return g.OutDegree(vertex)
}

// Traverses the set of edges in the graph, passing each edge to the
// provided closure.
func (g *Undirected) EachEdge(f func(edge Edge)) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	visited := set.NewNonTS()

	for source, adjacent := range g.list {
		for target, _ := range adjacent {
			e := &BaseEdge{U: source, V: target}
			if !visited.Has(e) {
				visited.Add(e)
				f(e)
			}
		}
	}
}

// Returns the density of the graph. Density is the ratio of edge count to the
// number of edges there would be in complete graph (maximum edge count).
func (g *Undirected) Density() float64 {
	g.mu.RLock()
	defer g.mu.RUnlock()

	order := g.Order()
	return 2 * float64(g.Size()) / float64(order*(order-1))
}

// Removes a vertex from the graph. Also removes any edges of which that
// vertex is a member.
func (g *Undirected) RemoveVertex(vertices ...Vertex) {
	if len(vertices) == 0 {
		return
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	for _, vertex := range vertices {
		if g.hasVertex(vertex) {
			f := func(adjacent Vertex) {
				delete(g.list[adjacent], vertex)
			}

			g.EachAdjacent(vertex, f)
			g.size = g.size - len(g.list[vertex])
			delete(g.list, vertex)
		}
	}
	return
}

// Adds a new edge to the graph.
func (g *Undirected) AddEdge(edge Edge) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.addEdge(edge)
}

// Adds a new edge to the graph.
func (g *Undirected) addEdge(edge Edge) (exists bool) {
	g.ensureVertex(edge.Source())
	g.ensureVertex(edge.Target())

	if _, exists = g.list[edge.Source()][edge.Target()]; !exists {
		g.list[edge.Source()][edge.Target()] = keyExists
		g.list[edge.Target()][edge.Source()] = keyExists
		g.size++
	}
	return !exists
}

// Removes an edge from the graph. This does NOT remove vertex members of the
// removed edge.
func (g *Undirected) RemoveEdge(edge Edge) {
	g.mu.Lock()
	defer g.mu.Unlock()

	s, t := edge.Both()
	if _, exists := g.list[s][t]; exists {
		delete(g.list[s], t)
		delete(g.list[t], s)
		g.size--
	}
}

