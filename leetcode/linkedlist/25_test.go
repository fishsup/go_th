package linkedlist

import (
	"testing"
)

/* 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。


示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
示例 2：


输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]
示例 3：

输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]
示例 4：

输入：head = [1], k = 1
输出：[1]
提示：

列表中节点的数量在范围 sz 内
1 <= sz <= 5000
0 <= Node.val <= 1000
1 <= k <= sz

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

/* *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* 剩余的节点需要保持原有顺序 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	} else {
		//新建一个头指针
		newHead := new(ListNode)
		var nextHead *ListNode
		var curHead = newHead
		curNode := head

		//初始化尾指针
		tailNode := curNode
		for i := 0; i < k-1; i++ {
			if tailNode != nil {
				tailNode = tailNode.Next
			}
		}

		for curNode != nil && tailNode != nil {
			for i := 0; i < k && curNode != nil; i++ {
				if i == 0 {
					nextHead = curNode
				}
				//头插法倒序
				a := curHead.Next
				curHead.Next = curNode
				curNode = curNode.Next
				curHead.Next.Next = a
				if tailNode != nil {
					tailNode = tailNode.Next
				}
			}
			curHead = nextHead
		}
		curHead.Next = curNode
		return newHead.Next
	}
}
func Test_reverseKGroup(t *testing.T) {
	head := &ListNode{Val: 1, Next: nil}
	var curHead = head
	for i := 2; i < 6; i++ {
		curHead.Next = &ListNode{Val: i, Next: nil}
		curHead = curHead.Next
	}
	t.Log(reverseKGroup(head, 3))
}
