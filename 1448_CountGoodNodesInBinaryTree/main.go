package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return traverse(root, math.MinInt)
}

func traverse(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}

	goodNode := 0
	if root.Val >= max {
		max = root.Val
		goodNode++
	}

	if root.Left != nil {
		goodNode += traverse(root.Left, max)
	}
	if root.Right != nil {
		goodNode += traverse(root.Right, max)
	}
	return goodNode
}
