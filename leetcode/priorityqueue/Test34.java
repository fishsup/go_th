package leetcode.priorityqueue;

import java.util.Comparator;
import java.util.PriorityQueue;

class Test34 {

    public ListNode mergeKLists(ListNode[] lists) {
        if(lists.length ==0){
            return null;
        }
        PriorityQueue<ListNode> queue = new PriorityQueue<>(lists.length, new Comparator<ListNode>() {
            @Override
            public int compare(ListNode o1, ListNode o2) {
                return o1.val -o2.val;
            }
        });
        for (ListNode node : lists) {
            if(node != null){
                queue.add(node);
            }
        }
        ListNode newHead = new ListNode();
        ListNode cur = newHead;
        while(!queue.isEmpty()){
            ListNode a = queue.poll();
            ListNode next  = a.next;
            a.next = null;

            cur.next = a;
            cur = cur.next;

            if(next!= null){
                queue.add(next);
            }
        }
        return newHead.next;
    }

    public class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }

}
