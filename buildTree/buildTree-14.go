package buildtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{}
	inmap := make(map[int]int)
	for i, n := range inorder {
		inmap[n] = i
	}
	traverse(preorder, inorder, root, 0, len(inorder)-1, 0, inmap)
	return root
}

func traverse(preorder []int, inorder []int, root *TreeNode, pindex, ileft, iright int, inmp map[int]int) *TreeNode {
	node := TreeNode{
		Val: preorder[pindex],
	}
	root = &node
	index := inmp[preorder[pindex]]

	left := inorder[ileft:index]
	right := inorder[index:iright]
	if len(left) == 1 {
		node.Left = &TreeNode{
			Val: left[0],
		}
	} else {
		traverse(preorder, inorder, &node, pindex, ileft+index, iright, inmp)
	}
	if len(right) == 1 {
		node.Right = &TreeNode{
			Val: right[0],
		}
	} else {
		traverse(preorder, inorder, &node, pindex, ileft, iright, inmp)
	}
}
