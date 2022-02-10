package linkedlist

//输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
// 示例1：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// 限制：
//
// 0 <= 链表长度 <= 1000
//
// 注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/
// Related Topics 递归 链表 👍 204 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
