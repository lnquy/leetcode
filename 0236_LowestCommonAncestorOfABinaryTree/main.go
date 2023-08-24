package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, found := traverse(root, p.Val, q.Val)
	return found
}

func traverse(node *TreeNode, a, b int) (c int, found *TreeNode) {
	if node == nil {
		return 0, nil
	}

	foundCounter := 0
	if node.Val == a || node.Val == b {
		foundCounter++
	}

	c, found = traverse(node.Left, a, b)
	if c >= 2 {
		return c, found
	}
	foundCounter += c
	c, found = traverse(node.Right, a, b)
	if c >= 2 {
		return c, found
	}
	foundCounter += c

	if foundCounter >= 2 {
		return foundCounter, node
	}
	return foundCounter, nil
}
