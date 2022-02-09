package _06

import (
	"fmt"
	"testing"
	"go_ph/leetcode/linkedlist"
)

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//
// Related Topics 递归 链表 👍 2264 👎 0

//leetcode submit region begin(Prohibit modification and deletion)


func (s *ListNode) String() string {
	return fmt.Sprintf("{Val: %v, Next: %v}", s.Val, s.Next)
}

func TestReverseList(t *testing.T) {
	head:= &ListNode{Val: 1}
	cur := head
	for i := 2; i < 10; i++ {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}
	fmt.Printf("%+v\n", head)
	fmt.Printf("%+v\n", reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	var newHead = new(ListNode)
	var cur = head
	for cur != nil {
		next := newHead.Next
		newHead.Next = cur
		cur = cur.Next
		newHead.Next.Next = next
	}
	return newHead.Next
}
