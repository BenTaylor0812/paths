package paths

import "fmt"

// Node describes a node structure for the shortest path algorithm.
type Node struct {
	Label    string
	NodeList map[string]int
}

// ShortestPath takes in a graph (array of nodes) and returns an array of the shortest path
func ShortestPath(graph []Node) {
	for _, node := range graph {
		fmt.Println(node)
	}
}
