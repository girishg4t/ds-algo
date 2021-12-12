/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package validbst

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	list := []int{}
	list = *dfs(root, &list)

	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode, list *[]int) *[]int {
	if node == nil {
		return list
	}
	dfs(node.Left, list)
	*list = append(*list, node.Val)
	dfs(node.Right, list)
	fmt.Println(list)
	return list
}
