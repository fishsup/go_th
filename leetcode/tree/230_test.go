package tree

import "testing"

//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,1,4,null,2], k = 1
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3
//
//
//
//
//
//
// 提示：
//
//
// 树中的节点数为 n 。
// 1 <= k <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
//
//
//
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 574 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TestName(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}
	t.Log(kthSmallest1(root, 3))
	t.Log(kthSmallest(root, 3))
}

func kthSmallest(root *TreeNode, k int) int {
	//中序遍历第k个
	var sk []*TreeNode
	for {
		for root != nil {
			sk = append(sk, root)
			root = root.Left
		}
		sk, root = sk[:len(sk)-1], sk[len(sk)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

var rtt []*TreeNode

func kthSmallest1(root *TreeNode, k int) int {
	//中序遍历第k个
	node := accessTree(root, k)
	return node.Val
}

func accessTree(tree *TreeNode, k int) *TreeNode {
	if tree.Left != nil {
		rt := accessTree(tree.Left, k)
		if rt != nil {
			return rt
		}
	}
	rtt = append(rtt, tree)
	if len(rtt) == k {
		return rtt[len(rtt)-1]
	}
	if tree.Right != nil {
		rt := accessTree(tree.Right, k)
		if rt != nil {
			return rt
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
