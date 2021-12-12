package maxdepth

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	level := 0
	return dfs(root, level)
}

func dfs(node *TreeNode, level int) int {
	if node == nil {
		return level
	}
	level++
	return int(math.Max(float64(dfs(node.Left, level)), float64(dfs(node.Right, level))))
}
