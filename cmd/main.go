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

	node1A := paths.Node{
		Label: "A",
		NodeList: map[string]int{
			"B": 1,
			"C": 6,
		},
	}

	node1B := paths.Node{
		Label: "B",
		NodeList: map[string]int{
			"A": 1,
			"C": 2,
			"D": 1,
		},
	}

	node1C := paths.Node{
		Label: "C",
		NodeList: map[string]int{
			"A": 6,
			"B": 2,
			"D": 2,
			"E": 5,
		},
	}

	node1D := paths.Node{
		Label: "D",
		NodeList: map[string]int{
			"B": 1,
			"C": 6,
			"E": 5,
		},
	}

	node1E := paths.Node{
		Label: "E",
		NodeList: map[string]int{
			"C": 5,
			"D": 5,
		},
	}

	graph := []*paths.Node{&nodeA, &nodeB, &nodeC, &nodeD, &nodeE, &nodeF}
	graph1 := []*paths.Node{&node1A, &node1B, &node1C, &node1D, &node1E}
	paths.TestMain(graph1)
	paths.TestMain(graph)

}
