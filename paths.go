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
func ShortestPath(startNode *Node, endNode *Node, graph []*Node) []string {
	var remainingNodes []*Node
	var cementedNodes []*Node

	startNode.previousNode = "start"

	endNode.nextNode = "end"

	currentNode := startNode

	shortestDistances := map[string]int{startNode.Label: 0}

	for len(cementedNodes) != len(graph) {
		fmt.Println(currentNode.Label)
		var minDist = shortestDistances[currentNode.Label]
		firstIt := true
		fmt.Println("========================")
		fmt.Println("Node is:", currentNode.Label)
		fmt.Println("========================")
		for k, v := range currentNode.NodeList {
			fmt.Printf("k value is : %s\n", k)
			node := findNode(k, graph)
			fmt.Println("Looking at node: ", node)
			if !find(k, remainingNodes) && !find(k, cementedNodes) {
				remainingNodes = append(remainingNodes, node)
				fmt.Println("The node label is:", node.Label)
				shortestDistances[node.Label] = minDist + v
				node.previousNode = currentNode.Label
				currentNode.nextNode = node.Label
			} else if find(k, remainingNodes) && !find(k, cementedNodes) {
				if minDist+v < shortestDistances[node.Label] {
					shortestDistances[node.Label] = minDist + v
					node.previousNode = currentNode.Label
					currentNode.nextNode = node.Label
				}
			} else {
				fmt.Println(k, "is a cemented element.")
			}
		}

		fmt.Println("Nodes left")
		for _, i := range remainingNodes {
			fmt.Println(*i)
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
		fmt.Println("The smallest node is:", smallestNode)
		currentNode = findNode(smallestNode, remainingNodes)
	}

	prevNode := "end"
	var nodeList []string
	for prevNode != "start" {
		var currentNode *Node
		if prevNode == "end" {
			currentNode = endNode
		} else {
			for _, i := range graph {
				if i.Label == prevNode {
					currentNode = i
				}
			}
		}
		nodeList = append([]string{currentNode.Label}, nodeList...)
		prevNode = currentNode.previousNode
	}

	return nodeList
}
