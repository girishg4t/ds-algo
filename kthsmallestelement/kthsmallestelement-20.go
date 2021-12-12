package kthsmallestelement

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	elements := []int{}
	dfs(root, &elements, k)
	return elements[k-1]
}

func dfs(root *TreeNode, elements *[]int, k int) {
	if root == nil {
		return
	}
	dfs(root.Left, elements, k)
	*elements = append(*elements, root.Val)
	if len(*elements)-1 == k {
		return
	}
	dfs(root.Right, elements, k)
}
