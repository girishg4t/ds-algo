package issametree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	firsttree := []int{}
	dfs(p, &firsttree)

	secondtree := []int{}
	dfs(q, &secondtree)
	if len(firsttree) != len(secondtree) {
		return false
	}
	fmt.Println(firsttree, secondtree)
	for i, n := range secondtree {
		if secondtree[i] != n {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, arr)
	*arr = append(*arr, node.Val)
	dfs(node.Right, arr)
}
