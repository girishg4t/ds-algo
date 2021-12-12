package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return clone(node, map[int]*Node{})
}

func clone(node *Node, hasmap map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, ok := hasmap[node.Val]; !ok {
		hasmap[node.Val] = &Node{Val: node.Val}
		for _, innernode := range node.Neighbors {
			hasmap[node.Val].Neighbors = append(hasmap[node.Val].Neighbors, clone(innernode, hasmap))
		}
	}

	return hasmap[node.Val]
}
