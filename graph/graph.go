package graph

import "fmt"

// Vertex defines the properties of a vertex
type Vertex struct {
	Name    string
	Visited bool
}

// Edge defines the properties of an edge
type Edge struct {
	Cost int
	From *Vertex
	To   *Vertex
}

// Graph defines the properties of a graph
type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

// New creates a new predefined graph
func New() Graph {
	va := &Vertex{Name: "A", Visited: false}
	vb := &Vertex{Name: "B", Visited: false}
	vc := &Vertex{Name: "C", Visited: false}
	vd := &Vertex{Name: "D", Visited: false}
	ve := &Vertex{Name: "E", Visited: false}
	vf := &Vertex{Name: "F", Visited: false}
	vg := &Vertex{Name: "G", Visited: false}

	g := Graph{Vertices: []*Vertex{va, vb, vc, vd, ve, vf, vg}}

	g.AddEdge(va, vb, 2)
	g.AddEdge(va, vc, 3)
	g.AddEdge(va, vd, 3)
	g.AddEdge(vb, vc, 4)
	g.AddEdge(vb, ve, 3)
	g.AddEdge(vc, vd, 5)
	g.AddEdge(vc, ve, 1)
	g.AddEdge(vd, vf, 7)
	g.AddEdge(ve, vf, 8)
	g.AddEdge(vf, vg, 9)

	return g
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to *Vertex, cost int) {
	e := &Edge{
		Cost: cost,
		From: from,
		To:   to,
	}

	g.Edges = append(g.Edges, e)
}

// PrintTree takes a tree as input and prints the origin, destination
// and cost of every edge
func PrintTree(tree []*Edge) {
	for _, edge := range tree {
		fmt.Printf("From: %s -> To: %s -> Cost: %d \n",
			edge.From.Name,
			edge.To.Name,
			edge.Cost,
		)
	}
}
