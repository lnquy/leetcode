package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	return max(traverse(root.Left, true, 0), traverse(root.Right, false, 0))
}

func traverse(root *TreeNode, isLeft bool, length int) int {
	if root == nil {
		return length
	}

	if isLeft {
		return max(traverse(root.Left, true, 0), traverse(root.Right, false, length+1))
	}
	return max(traverse(root.Left, true, length+1), traverse(root.Right, false, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
