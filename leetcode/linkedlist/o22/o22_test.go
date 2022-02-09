package o22

//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
// 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//
//
//
// 示例：
//
//
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.
// Related Topics 链表 双指针 👍 323 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k <= 0 || head == nil {
		panic("error input")
	}
	tail := head
	for i := 0; i < k-1; i++ {
		if tail.Next == nil {
			panic("error list length")
		}
		tail = tail.Next
	}
	for tail.Next != nil {
		tail = tail.Next
		head = head.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
