package main

import (
	"github.com/BenTaylor0812/paths"
)

func main() {
	nodeA := paths.Node{
		Label: "A",
		NodeList: map[string]int{
			"B": 4,
			"C": 2,
		},
	}

	nodeB := paths.Node{
		Label: "B",
		NodeList: map[string]int{
			"D": 10,
			"C": 5,
		},
	}

	nodeC := paths.Node{
		Label: "C",
		NodeList: map[string]int{
			"E": 3,
		},
	}

	nodeD := paths.Node{
		Label: "D",
		NodeList: map[string]int{
			"F": 11,
		},
	}

	nodeE := paths.Node{
		Label: "E",
		NodeList: map[string]int{
			"D": 4,
		},
	}

	nodeF := paths.Node{
		Label:    "F",
		NodeList: map[string]int{},
	}

	graph := []paths.Node{nodeA, nodeB, nodeC, nodeD, nodeE, nodeF}
	paths.ShortestPath(graph)
}
