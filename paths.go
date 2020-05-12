package paths

import (
	"fmt"
)

// Node describes a node structure for the shortest path algorithm.
type Node struct {
	Label        string
	NodeList     map[string]int
	distance     int
	previousNode string
	nextNode     string
}

func deleteElement(nodes []*Node, node Node) []*Node {
	for i, v := range nodes {
		if v.Label == node.Label {
			nodes = append(nodes[:i], nodes[i+1:]...)
		}
	}
	return nodes
}

func obtainNode(nodes []*Node, label string) *Node {
	var rNode *Node
	for _, node := range nodes {
		if node.Label == label {
			rNode = node
		}
	}
	return rNode
}

func customCopy(copyTo []*Node, copyFrom []*Node) {
	for i, v := range copyFrom {
		copyTo[i] = v
	}
}

func find(node string, nodes []*Node) bool {
	var rVal bool
	for _, i := range nodes {
		if node == i.Label {
			rVal = true
		}
	}
	return rVal
}

func findNode(node string, nodes []*Node) *Node {
	var rNode *Node
	for _, i := range nodes {
		if i.Label == node {
			rNode = i
			break
		}
	}
	return rNode
}

// ShortestPath takes in a graph (array of nodes) and returns an array of the shortest path
func ShortestPath(graph []*Node) {
	var remainingNodes []*Node
	//customCopy(remainingNodes, graph)
	var cementedNode []*Node

	startNode := graph[0]
	startNode.distance = 0
	startNode.previousNode = "first"

	endNode := graph[len(graph)-1]
	endNode.nextNode = "end"

	currentNode := startNode
	// cementedNode = append(cementedNode, startNode)
	//remainingNodes = deleteElement(remainingNodes, *startNode)

	shortestDistances := map[string]int{startNode.Label: 0}

	for len(cementedNode) != len(graph) {
		var minDist = shortestDistances[currentNode.Label]
		firstIt := true
		for k, v := range currentNode.NodeList {
			node := findNode(k, graph)
			if !find(k, remainingNodes) && !find(k, cementedNode) {
				remainingNodes = append(remainingNodes, node)
				shortestDistances[node.Label] = minDist + v
			} else if find(k, remainingNodes) && !find(k, cementedNode) {
				if minDist+v < shortestDistances[node.Label] {
					shortestDistances[node.Label] = minDist + v
					node.previousNode = currentNode.Label
					currentNode.nextNode = node.Label
				}
			}

			if v < minDist || firstIt {
				minDist = v
				currentNode.nextNode = k
				firstIt = false
			}
		}

		cementedNode = append(cementedNode, currentNode)
		remainingNodes = deleteElement(remainingNodes, *currentNode)

		var smallest int
		var smallestNode string
		firstIt = true
		for k, v := range shortestDistances {
			if v < smallest || firstIt {
				smallest = v
				smallestNode = k
			}
		}
		fmt.Println(smallestNode)
		currentNode = findNode(smallestNode, remainingNodes)

		fmt.Printf("Cemented nodes: ")
		for _, i := range cementedNode {
			fmt.Printf(" [%v], ", *i)
		}
		fmt.Printf("\n")
		fmt.Printf("Remaining nodes ")
		for _, i := range remainingNodes {
			fmt.Printf(" [%v], ", *i)
		}
		fmt.Printf("\n")
		fmt.Println(shortestDistances)

		// for k := range shortestDistances {
		// 	if find(k, remainingNodes) {

		// 	}
		// }
	}
	// fmt.Println(shortestDistances)
	// for _, i := range cementedNode {
	// 	fmt.Println(*i)
	// }

	// var minDist int
	// firstIt := true

	// // fmt.Printf("Closest node is: %s, with a distance of: %d.", currentNode.nextNode, minDist)
	// minDist = currentNode.distance + minDist
	// currentNode = obtainNode(graph, currentNode.nextNode)
	// cementedNode = append(cementedNode, currentNode)
	// currentNode.distance = minDist
	// shortestDistances[currentNode.Label] = currentNode.distance
	// fmt.Println(shortestDistances)

	// for currentNode.nextNode != "end" {
	// 	fmt.Println(currentNode)
	// 	firstIt = true
	// 	for k, v := range currentNode.NodeList {
	// 		if v < minDist || firstIt {
	// 			minDist = v
	// 			currentNode.nextNode = k
	// 			firstIt = false
	// 		}
	// 	}
	// 	// fmt.Printf("Closest node is: %s, with a distance of: %d.", currentNode.nextNode, minDist)
	// 	minDist = currentNode.distance + minDist
	// 	fmt.Println(currentNode)
	// 	currentNode = obtainNode(graph, currentNode.nextNode)
	// 	fmt.Println(currentNode)
	// 	fmt.Println(endNode)
	// 	cementedNode = append(cementedNode, currentNode)
	// 	currentNode.distance = minDist
	// 	shortestDistances[currentNode.Label] = currentNode.distance
	// 	if currentNode.Label == "F" {
	// 		break
	// 	}
	// }

	// fmt.Println(shortestDistances)
}

// TestMain -
func TestMain(graph []*Node) {
	ShortestPath(graph)
}
