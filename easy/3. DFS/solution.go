package program

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// first add the node.name to the array
	array = append(array, n.Name)
	for _, child_node := range n.Children {
		array = child_node.DepthFirstSearch(array)
	}
	return array
}
