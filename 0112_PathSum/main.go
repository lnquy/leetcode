package main

func main() {
	// No test was written for this exercise since I couldn't guarantee the structure of binary tree after inserted.
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
