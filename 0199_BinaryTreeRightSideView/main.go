package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) (res []int) {
	rightestNodePerLevel := make(map[int]int)
	groupNodesByLevel(root, 0, rightestNodePerLevel)
	for i := 0; i < len(rightestNodePerLevel); i++ {
		if v, ok := rightestNodePerLevel[i]; ok {
			res = append(res, v)
		}
	}
	return res
}

func groupNodesByLevel(node *TreeNode, level int, rightestNodePerLevel map[int]int) {
	if node == nil {
		return
	}

	rightestNodePerLevel[level] = node.Val
	groupNodesByLevel(node.Left, level+1, rightestNodePerLevel)
	groupNodesByLevel(node.Right, level+1, rightestNodePerLevel)
	return
}
