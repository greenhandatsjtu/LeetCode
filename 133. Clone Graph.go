package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var clone func(node *Node) *Node
	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := visited[node]; ok {
			return n
		}
		newNode := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
		visited[node] = newNode
		for i, n := range node.Neighbors {
			newNode.Neighbors[i] = clone(n)
		}
		return newNode
	}
	return clone(node)
}
