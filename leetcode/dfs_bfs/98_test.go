package dfsbfs

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dbf(root.Left, math.MinInt, root.Val) && dbf(root.Right, root.Val, math.MaxInt)
}

func dbf(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return dbf(node.Left, min, node.Val) && dbf(node.Right, node.Val, max)
}
