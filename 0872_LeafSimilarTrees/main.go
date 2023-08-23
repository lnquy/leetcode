package main

import "reflect"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(toArray(root1, []int{}), toArray(root2, []int{}))
}

func toArray(root *TreeNode, dstArr []int) []int {
	if root.Left != nil {
		dstArr = toArray(root.Left, dstArr)
	}
	if root.Left == nil && root.Right == nil {
		dstArr = append(dstArr, root.Val)
	}
	if root.Right != nil {
		dstArr = toArray(root.Right, dstArr)
	}
	return dstArr
}
