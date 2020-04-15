package main

import (
	"algorithms/graph"
	"algorithms/spanningtrees"
	"fmt"
)

func main() {
	fmt.Print("1. Kruskal")
	fmt.Print("Select an algorithm to run: ")
	var option int
	fmt.Scan(&option)

	switch option {
	case 1:
		g := graph.New()
		graph.PrintTree(spanningtrees.Kruskal(&g))
	}

}
