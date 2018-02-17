package graph

type DirectedGraph struct {
	// base class is graph
	graph
}

func NewDirectedGraph() *DirectedGraph {
	g := DirectedGraph{}
	g.edges = make(map[VertexId]map[VertexId]int)
	g.edgesCount = 0
	g.isDirected = true
	return &g
}
