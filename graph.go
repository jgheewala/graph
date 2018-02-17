package graph

import "errors"

type VertexId uint32
type Vertices []VertexId

type Edge struct {
	From VertexId
	To   VertexId
}

// graph basicd data-strucute
type graph struct {
	// all edges in the graph
	edges map[VertexId]map[VertexId]int
	// total edgeCount
	edgesCount int
	// check wheter the graph is directed or not
	isDirected bool
}

// add a new vertex v to graph edge
func (g *graph) AddVertex(v VertexId) error {
	edge, exists := g.edges[v]
	if exists || edge != nil {
		return errors.New("Vertex already exists")
	}
	g.edges[v] = make(map[VertexId]int)
	return nil
}

// remove a vertex v from the graph edge
func (g *graph) RemoveVertex(v VertexId) error {
	if !g.IsVertex(v) {
		return errors.New("No vertex exists and hence cannot do remove")
	}

	delete(g.edges, v)
	for _, connectedVertices := range g.edges {
		delete(connectedVertices, v)
	}
	return nil
}

func (g *graph) IsVertex(v VertexId) bool {
	_, exists := g.edges[v]
	return exists
}

func (g *graph) AddEdge(from, to VertexId, weight int) error {
	if from == to {
		return errors.New("Cannot add self edge...aka loop")
	}
	if !g.IsVertex(from) || !g.IsVertex(to) {
		return errors.New("Vertices doesn't exists")
	}
	ft, _ := g.edges[from][to]
	tf, _ := g.edges[to][from]

	if ft > 0 || tf > 0 {
		return errors.New("Edges already defined")
	}
	g.TouchVertex(from)
	g.TouchVertex(to)
	g.edges[from][to] = weight
	if !g.isDirected {
		g.edges[to][from] = weight

	}
	g.edgesCount++
	return nil
}

func (g *graph) RemoveEdge(from, to VertexId) error {
	ft, _ := g.edges[from][to]
	tf, _ := g.edges[to][from]

	if ft == -1 || tf == -1 {
		return errors.New("Edge doesn't exists")
	}

	g.edges[from][to] = -1

	if !g.isDirected {
		g.edges[to][from] = -1
	}

	g.edgesCount--
	return nil
}

func (g *graph) IsEdge(from, to VertexId) bool {
	connected, exists := g.edges[from]
	if !exists {
		return exists
	}

	weight := connected[to]
	return weight > 0
}

func (g *graph) TouchVertex(v VertexId) {
	if _, ok := g.edges[v]; !ok {
		g.edges[v] = make(map[VertexId]int)
	}
}

type EdgesIterable interface {
	EdgesIter() <-chan Edge
}

type VerticesIterable interface {
	VerticesIter() <-chan VertexId
}

func (g *graph) EdgesIter() <-chan Edge {
	ch := make(chan Edge)
	go func() {
		for from, connectedVertices := range g.edges {
			for to, _ := range connectedVertices {
				if g.isDirected {
					ch <- Edge{from, to}
				} else {
					if from < to {
						ch <- Edge{from, to}
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (g *graph) VerticesIter() <-chan VertexId {
	ch := make(chan VertexId)
	go func() {
		for vertex, _ := range g.edges {
			ch <- vertex
		}
		close(ch)
	}()
	return ch
}

func (g *graph) Order() int {
	return len(g.edges)
}

func (g *graph) VerticesCount() int {
	return len(g.edges)
}

func (g *graph) EdgesCount() int {
	return g.edgesCount
}

func (g *graph) GetEdge(from, to VertexId) int {
	return g.edges[from][to]
}
