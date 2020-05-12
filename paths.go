package paths

import "fmt"

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
	var cementedNodes []*Node

	startNode := graph[0]
	startNode.previousNode = "first"

	endNode := graph[len(graph)-1]
	endNode.nextNode = "end"

	currentNode := startNode

	shortestDistances := map[string]int{startNode.Label: 0}

	for len(cementedNodes) != len(graph) {
		var minDist = shortestDistances[currentNode.Label]
		firstIt := true
		for k, v := range currentNode.NodeList {
			node := findNode(k, graph)
			if !find(k, remainingNodes) && !find(k, cementedNodes) {
				remainingNodes = append(remainingNodes, node)
				shortestDistances[node.Label] = minDist + v
				node.previousNode = currentNode.Label
				currentNode.nextNode = node.Label
			} else if find(k, remainingNodes) && !find(k, cementedNodes) {
				if minDist+v < shortestDistances[node.Label] {
					shortestDistances[node.Label] = minDist + v
					node.previousNode = currentNode.Label
					currentNode.nextNode = node.Label
				}
			}
		}

		cementedNodes = append(cementedNodes, currentNode)
		remainingNodes = deleteElement(remainingNodes, *currentNode)

		var smallest int
		var smallestNode string
		firstIt = true
		for k, v := range shortestDistances {
			if (v < smallest || firstIt) && find(k, remainingNodes) {
				smallest = v
				smallestNode = k
				firstIt = false
			}
		}

		currentNode = findNode(smallestNode, remainingNodes)
	}

	nextNode := "start"
	var nodeList []string
	for nextNode != "end" {
		var currentNode *Node
		if nextNode == "start" {
			currentNode = startNode
		} else {
			for _, i := range graph {
				if i.Label == nextNode {
					currentNode = i
				}
			}
		}
		nodeList = append(nodeList, currentNode.Label)
		nextNode = currentNode.nextNode
	}
	fmt.Println(nodeList)
}

// TestMain -
func TestMain(graph []*Node) {
	ShortestPath(graph)
}
