package tree

import "testing"

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 广度优先搜索 二叉树 👍 1183 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Test_Slice(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}
	t.Log(levelOrder(root))

}

func levelOrder(root *TreeNode) [][]int {
	lstq := []*TreeNode{}
	if root != nil {
		lstq = append(lstq, root)
	}

	rs := [][]int{}
	i := 0
	curq := []*TreeNode{}

	for len(lstq) != 0 {
		if len(rs) < i+1 {
			rs = append(rs, []int{})
		}
		node := lstq[0]
		lstq = lstq[1:]
		rs[i] = append(rs[i], node.Val)
		if node.Left != nil {
			curq = append(curq, node.Left)
		}
		if node.Right != nil {
			curq = append(curq, node.Right)
		}

		if len(lstq) == 0 {
			lstq = curq
			curq = []*TreeNode{}
			i++
		}

		
	}
	return rs
}

//leetcode submit region end(Prohibit modification and deletion)
