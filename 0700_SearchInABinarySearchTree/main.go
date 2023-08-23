package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == val {
		return root
	}
	if root.Left != nil {
		if r := searchBST(root.Left, val); r != nil {
			return r
		}
	}
	if r := searchBST(root.Right, val); r != nil {
		return r
	}
	return nil
}
