package main

import (
	"fmt"
	"sort"
)

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

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to *Vertex, cost int) {
	e := &Edge{
		Cost: cost,
		From: from,
		To:   to,
	}

	g.Edges = append(g.Edges, e)
}

// Kruskal implements the Kruskal algorithm to find the spanning tree of the graph g.
// It sorts the edges of the graph by cost in ascending order.
// Then it iterates all of the Edges and if the destination edge hasn't been visited yet
// it's added to the tree and both vertices connected by the edge are set as visited.
func (g *Graph) Kruskal() []*Edge {
	sort.Slice(g.Edges, func(i, j int) bool { return g.Edges[i].Cost < g.Edges[j].Cost })

	var tree []*Edge
	for _, edge := range g.Edges {
		if !edge.To.Visited {
			edge.From.Visited = true
			edge.To.Visited = true
			tree = append(tree, edge)
		}
	}

	return tree
}

func printTree(tree []*Edge) {
	for _, edge := range tree {
		fmt.Printf("From: %s -> To: %s -> Cost: %d \n",
			edge.From.Name,
			edge.To.Name,
			edge.Cost,
		)
	}
}

func main() {
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

	tree := g.Kruskal()
	printTree(tree)

}
