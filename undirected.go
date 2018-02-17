package graph

type UnDirectedGraph struct {
	graph
}

func NewUnDirectedGraph() *UnDirectedGraph {
	g := UnDirectedGraph{}
	g.edges = make(map[VertexId]map[VertexId]int)
	g.edgesCount = 0
	g.isDirected = false
	return &g
}
