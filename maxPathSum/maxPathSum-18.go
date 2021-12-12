package maxpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxValue := []int{math.MinInt}
	traverse(root, maxValue)
	return maxValue[0]
}

func traverse(root *TreeNode, maxValue []int) int {
	if root == nil {
		return 0
	}

	leftMax := int(math.Max(float64(traverse(root.Left, maxValue)), float64(0)))
	rightMax := int(math.Max(float64(traverse(root.Right, maxValue)), float64(0)))

	maxValue[0] = int(math.Max(float64(maxValue[0]), float64(root.Val+leftMax+rightMax)))

	return root.Val + int(math.Max(float64(leftMax), float64(rightMax)))
}
