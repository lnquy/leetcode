package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	sumMap := make(map[int]int, 1000)
	sumMap[0] = 1
	return traverse(root, 0, targetSum, sumMap)
}

func traverse(root *TreeNode, currSum, targetSum int, sumMap map[int]int) (foundPaths int) {
	if root == nil {
		return 0
	}

	currSum += root.Val
	if cnt, ok := sumMap[currSum-targetSum]; ok {
		foundPaths += cnt
	}
	if _, ok := sumMap[currSum]; ok {
		sumMap[currSum]++
	} else {
		sumMap[currSum] = 1
	}

	foundPaths += traverse(root.Left, currSum, targetSum, sumMap)
	foundPaths += traverse(root.Right, currSum, targetSum, sumMap)

	sumMap[currSum]--
	return foundPaths
}
