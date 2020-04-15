package spanningtrees

import (
	"algorithms/graph"
	"sort"
)

// Kruskal implements the Kruskal algorithm to find the spanning tree of the graph g.
// It sorts the edges of the graph by cost in ascending order.
// Then it iterates all of the Edges and if the destination edge hasn't been visited yet
// it's added to the tree and both vertices connected by the edge are set as visited.
func Kruskal(g *graph.Graph) []*graph.Edge {
	sort.Slice(g.Edges, func(i, j int) bool { return g.Edges[i].Cost < g.Edges[j].Cost })

	var tree []*graph.Edge
	for _, edge := range g.Edges {
		if !edge.To.Visited {
			edge.From.Visited = true
			edge.To.Visited = true
			tree = append(tree, edge)
		}
	}

	return tree
}
