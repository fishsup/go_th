package tree

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[2,1]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1277 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	rs := []int{}
	q := []*TreeNode{}
	for root != nil {
		q = append(q, root)
		root = root.Left
	}
	for len(q) != 0 {
		cur := q[len(q)-1]
		q = q[:len(q)-1]
		rs = append(rs, cur.Val)
		if cur.Right != nil {
			a := cur.Right
			for a != nil {
				q = append(q, a)
				a = a.Left
			}
		}
	}
	return rs
}

/* 

// 递归实现
var rs []int

func inorderTraversal1(root *TreeNode) []int {
	rs = []int{}
	if root != nil {
		access(root)
	}
	return rs
}

func access(node *TreeNode) {
	if node != nil {
		if node.Left != nil {
			access(node.Left)
		}
		rs = append(rs, node.Val)
		if node.Right != nil {
			access(node.Right)
		}
	}
}
*/

//leetcode submit region end(Prohibit modification and deletion)
