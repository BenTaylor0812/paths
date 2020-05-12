package main

import (
	"fmt"

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

	gameA1 := paths.Node{
		Label: "A1",
		NodeList: map[string]int{
			"A2": 1,
			"B1": 1,
		},
	}

	gameA2 := paths.Node{
		Label: "A2",
		NodeList: map[string]int{
			"A1": 1,
			"A3": 1,
		},
	}

	gameA3 := paths.Node{
		Label: "A3",
		NodeList: map[string]int{
			"A2": 1,
			"A4": 1,
		},
	}

	gameA4 := paths.Node{
		Label: "A4",
		NodeList: map[string]int{
			"A3": 1,
			"A5": 1,
		},
	}

	gameA5 := paths.Node{
		Label: "A5",
		NodeList: map[string]int{
			"A4": 1,
			"B5": 1,
			"A6": 1,
		},
	}

	gameA6 := paths.Node{
		Label: "A6",
		NodeList: map[string]int{
			"A5": 1,
			"B6": 1,
		},
	}

	gameB1 := paths.Node{
		Label: "B1",
		NodeList: map[string]int{
			"A1": 1,
			"C1": 1,
		},
	}

	gameB2 := paths.Node{
		Label: "B2",
		NodeList: map[string]int{
			"C2": 1,
			"B3": 1,
		},
	}

	gameB3 := paths.Node{
		Label: "B3",
		NodeList: map[string]int{
			"B2": 1,
			"C3": 1,
		},
	}

	gameB4 := paths.Node{
		Label: "B4",
		NodeList: map[string]int{
			"C4": 1,
		},
	}

	gameB5 := paths.Node{
		Label: "B5",
		NodeList: map[string]int{
			"A5": 1,
			"C5": 1,
		},
	}

	gameB6 := paths.Node{
		Label: "B6",
		NodeList: map[string]int{
			"A6": 1,
		},
	}

	gameC1 := paths.Node{
		Label: "C1",
		NodeList: map[string]int{
			"B1": 1,
			"C2": 1,
		},
	}

	gameC2 := paths.Node{
		Label: "C2",
		NodeList: map[string]int{
			"C1": 1,
			"B2": 1,
		},
	}

	gameC3 := paths.Node{
		Label: "C3",
		NodeList: map[string]int{
			"B3": 1,
			"C4": 1,
		},
	}

	gameC4 := paths.Node{
		Label: "C4",
		NodeList: map[string]int{
			"C3": 1,
			"D4": 1,
			"B4": 1,
		},
	}

	gameC5 := paths.Node{
		Label: "C5",
		NodeList: map[string]int{
			"B5": 1,
			"C6": 1,
		},
	}

	gameC6 := paths.Node{
		Label: "C6",
		NodeList: map[string]int{
			"C5": 1,
			"D6": 1,
		},
	}

	gameD1 := paths.Node{
		Label: "D1",
		NodeList: map[string]int{
			"E1": 1,
			"D2": 1,
		},
	}

	gameD2 := paths.Node{
		Label: "D2",
		NodeList: map[string]int{
			"D1": 1,
			"E2": 1,
		},
	}

	gameD3 := paths.Node{
		Label: "D3",
		NodeList: map[string]int{
			"E3": 1,
			"D4": 1,
		},
	}

	gameD4 := paths.Node{
		Label: "D4",
		NodeList: map[string]int{
			"D3": 1,
			"C4": 1,
		},
	}

	gameD5 := paths.Node{
		Label: "D5",
		NodeList: map[string]int{
			"E5": 1,
			"D6": 1,
		},
	}

	gameD6 := paths.Node{
		Label: "D6",
		NodeList: map[string]int{
			"D5": 1,
			"C6": 1,
		},
	}

	gameE1 := paths.Node{
		Label: "E1",
		NodeList: map[string]int{
			"D1": 1,
			"F1": 1,
		},
	}

	gameE2 := paths.Node{
		Label: "E2",
		NodeList: map[string]int{
			"D2": 1,
			"F2": 1,
		},
	}

	gameE3 := paths.Node{
		Label: "E3",
		NodeList: map[string]int{
			"D3": 1,
			"F3": 1,
		},
	}

	gameE4 := paths.Node{
		Label: "E4",
		NodeList: map[string]int{
			"F4": 1,
		},
	}

	gameE5 := paths.Node{
		Label: "E5",
		NodeList: map[string]int{
			"D5": 1,
		},
	}

	gameE6 := paths.Node{
		Label: "E6",
		NodeList: map[string]int{
			"F6": 1,
		},
	}

	gameF1 := paths.Node{
		Label: "F1",
		NodeList: map[string]int{
			"E1": 1,
		},
	}

	gameF2 := paths.Node{
		Label: "F2",
		NodeList: map[string]int{
			"E2": 1,
			"F3": 1,
		},
	}

	gameF3 := paths.Node{
		Label: "F3",
		NodeList: map[string]int{
			"F2": 1,
			"E3": 1,
			"F4": 1,
		},
	}

	gameF4 := paths.Node{
		Label: "F4",
		NodeList: map[string]int{
			"F3": 1,
			"F5": 1,
			"E4": 1,
		},
	}

	gameF5 := paths.Node{
		Label: "F5",
		NodeList: map[string]int{
			"F4": 1,
			"F6": 1,
		},
	}

	gameF6 := paths.Node{
		Label: "F6",
		NodeList: map[string]int{
			"F5": 1,
			"E6": 1,
		},
	}

	graph := []*paths.Node{&nodeA, &nodeB, &nodeC, &nodeD, &nodeE, &nodeF}
	graph1 := []*paths.Node{&node1A, &node1B, &node1C, &node1D, &node1E}
	game := []*paths.Node{
		&gameA1, &gameA2, &gameA3, &gameA4, &gameA5, &gameA6,
		&gameB1, &gameB2, &gameB3, &gameB4, &gameB5, &gameB6,
		&gameC1, &gameC2, &gameC3, &gameC4, &gameC5, &gameC6,
		&gameD1, &gameD2, &gameD3, &gameD4, &gameD5, &gameD6,
		&gameE1, &gameE2, &gameE3, &gameE4, &gameE5, &gameE6,
		&gameF1, &gameF2, &gameF3, &gameF4, &gameF5, &gameF6,
	}

	if false {
		fmt.Println(paths.ShortestPath(&node1B, &node1C, graph1))
		fmt.Println(paths.ShortestPath(&nodeA, &nodeF, graph))
	}
	fmt.Println(paths.ShortestPath(&gameE6, &gameC6, game))
}
