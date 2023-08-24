package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	sumPerLevel := make(map[int]int)

	getSumPerLevel(root, 1, sumPerLevel)
	max := math.MinInt
	maxLevel := math.MaxInt
	for k, v := range sumPerLevel {
		if max < v {
			maxLevel = k
			max = v
		} else if max == v && k < maxLevel {
			maxLevel = k
		}
	}
	return maxLevel
}

func getSumPerLevel(node *TreeNode, level int, sumPerLevel map[int]int) {
	if node == nil {
		return
	}

	sumPerLevel[level] += node.Val
	getSumPerLevel(node.Left, level+1, sumPerLevel)
	getSumPerLevel(node.Right, level+1, sumPerLevel)
}
