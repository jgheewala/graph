package graph

import (
	"fmt"
	"testing"
)

var (
	testUndirectedGraphVertices = 10
)

func TestUnDirectedGraph(t *testing.T) {
	g := NewUnDirectedGraph()
	for idx := 0; idx < testUndirectedGraphVertices; idx++ {
		v := VertexId(idx)
		g.AddVertex(v)
	}

	if len(g.edges) != testUndirectedGraphVertices {
		t.Error("Failed to add Vertices to graph expected count:", testUndirectedGraphVertices,
			"received count", len(g.edges))
		return
	}

	t.Log(g)
	for idx := 0; idx < testUndirectedGraphVertices; idx++ {
		g.AddEdge(VertexId(idx), VertexId(idx%2), 1)

	}
	t.Log(g)
	if g.IsEdge(0, 8) == false || g.IsEdge(0, 9) == true || g.IsVertex(2) != true {
		fmt.Println(g)
		t.Error()
		return
	}
	// AddEdge should fail for already existing Edge
	err := g.AddEdge(0, 2, 1)
	if err == nil {
		fmt.Println(g)
		t.Error()
	}

	// AddVertex should fail for already existing vertex
	err = g.AddVertex(0)
	if err == nil {
		fmt.Println(g)
		t.Error()
	}

	g.RemoveVertex(VertexId(9))

	if g.IsVertex(VertexId(9)) {
		fmt.Println(g.edges[9] == nil)
		t.Error()
	}

	// RemoveVertex should fail for unknown vertex
	err = g.RemoveVertex(VertexId(9))

	if err == nil {
		fmt.Println(g.edges[9] == nil)
		t.Error()
	}

	g.RemoveEdge(0, 8)

	if g.IsEdge(VertexId(0), VertexId(8)) == true || g.edgesCount != 7 {
		fmt.Println(g.IsEdge(VertexId(0), VertexId(8)), g.edgesCount)
		t.Error()
	}
	// RemoveEdge should fail for unknown egde
	err = g.RemoveEdge(0, 8)

	if err == nil {
		fmt.Println(g)
		t.Error()
	}
	c := g.EdgesIter()

	countEdge := 0
	for _ = range c {
		countEdge++
	}

	if g.EdgesCount() != countEdge {
		t.Error()
	}

	d := g.VerticesIter()
	verticesCount := g.Order()
	countVertices := 0

	for _ = range d {
		countVertices++
	}

	if countVertices != verticesCount {
		fmt.Println(countVertices, g.edges)
		t.Error()
	}

	g.TouchVertex(9)
	if _, ok := g.edges[9]; !ok {
		t.Error()
	}
}
