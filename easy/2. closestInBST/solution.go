package program

import "math"

type BST struct {
	value int

	left  *BST
	right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	return tree.FindClosestValueBST(target, tree.value)
}
func (tree *BST) FindClosestValueBST(target, closest int) int {
	node := tree
	for node != nil {
		if diff(target, closest) > diff(target, node.value) {
			closest = node.value
		}
		if target < node.value {
			node = node.left
		} else if target > node.value {
			node = node.right
		} else {
			break
		}
	}
	return closest

}

func diff(a, b int) int {
	result := math.Abs(float64(a) - float64(b))
	return int(result)
}
