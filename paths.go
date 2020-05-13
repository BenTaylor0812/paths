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
}

// This function deleted an element node from the slice nodes
func deleteElement(nodes []*Node, node Node) []*Node {
	for i, v := range nodes {
		if v.Label == node.Label {
			nodes = append(nodes[:i], nodes[i+1:]...)
		}
	}
	return nodes
}

// This function find a node corresponding to a label and returns the node
func obtainNode(nodes []*Node, label string) *Node {
	var rNode *Node
	for _, node := range nodes {
		if node.Label == label {
			rNode = node
		}
	}
	return rNode
}

// This function copies copyFrom to copyTo to avoid assignment issues
func customCopy(copyTo []*Node, copyFrom []*Node) {
	for i, v := range copyFrom {
		copyTo[i] = v
	}
}

// This finds if a node with label node is in the nodes slice, may be able to fuse with findNode
func find(node string, nodes []*Node) bool {
	var rVal bool
	for _, i := range nodes {
		if node == i.Label {
			rVal = true
		}
	}
	return rVal
}

// Similar to find but returns a node and an error
func findNode(node string, nodes []*Node) (*Node, error) {
	var rNode *Node
	var found bool
	for _, i := range nodes {
		if i.Label == node {
			rNode = i
			found = true
			break
		}
	}
	if !found {
		return rNode, fmt.Errorf("node not found")
	}
	return rNode, nil
}

// ShortestPath takes in a graph (array of nodes) and returns an array of the shortest path
func ShortestPath(startNode *Node, endNode *Node, graph []*Node) ([]string, error) {
	var remainingNodes []*Node
	var cementedNodes []*Node

	startNode.previousNode = "start"

	currentNode := startNode

	shortestDistances := map[string]int{startNode.Label: 0}

	// For a connected graph we want all the nodes to be considered, after a node is considered it is "cemented",
	// this means we have found the optimal way to get to that node from the starting node.
	for len(cementedNodes) != len(graph) {
		var minDist = shortestDistances[currentNode.Label]
		firstIt := true

		// k is the label of the nodes that currentNode is connected to and v is the distance to said node
		for k, v := range currentNode.NodeList {
			node, err := findNode(k, graph)
			if err != nil {
				fmt.Printf("%s is not found within the graph.\n", k)
				return []string{""}, fmt.Errorf("a node was not found")
			}

			// If the k node is not in the remaining nodes or the cemented nodes, that means it hasn't been considered
			// and isn't on the list to be. SO we add it to the remaining nodes to be processed later.
			// If it is in the queue, then we check that it if it has a shorter path using the current node, and if it does,
			// update the order of nodes.
			if !find(k, remainingNodes) && !find(k, cementedNodes) {
				remainingNodes = append(remainingNodes, node)
				shortestDistances[node.Label] = minDist + v
				node.previousNode = currentNode.Label
			} else if find(k, remainingNodes) && !find(k, cementedNodes) {
				if minDist+v < shortestDistances[node.Label] {
					shortestDistances[node.Label] = minDist + v
					node.previousNode = currentNode.Label
				}
			} else {
			}
		}

		cementedNodes = append(cementedNodes, currentNode)
		remainingNodes = deleteElement(remainingNodes, *currentNode)

		// Here we check for which node in remaining nodes has the shortest distance from the startingNode, and we set that
		// one to be the next node to be considered.
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

		var err error
		currentNode, err = findNode(smallestNode, remainingNodes)
		if err != nil && len(cementedNodes) != len(graph) {
			fmt.Printf("%s is not found within the graph. Can't set currentNode.\n", smallestNode)
			return []string{""}, fmt.Errorf("a node was not found")
		}
	}

	// When all nodes have been considered, we can provide the path it takes by looking backwards at the previous
	// node from each node, starting at the final node.
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
	fmt.Println(shortestDistances)

	// Might be an idea to update the retun values to provide the distance from start to end.
	return nodeList, nil
}
