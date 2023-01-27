package pq;
import java.util.PriorityQueue;   
public class QueueP{
    public static ListNode mergeKLists(ListNode[] lists) {
        ListNode dummy = new ListNode(-1), tail = dummy;
        PriorityQueue<ListNode> q = new PriorityQueue<>((a,b)->a.val-b.val);
        for (ListNode node : lists) {
            if (node != null) q.add(node);
        }
        while (!q.isEmpty()) {
            ListNode poll = q.poll();
            tail.next = poll;
            tail = tail.next;
            if (poll.next != null) q.add(poll.next);
        }
        return dummy.next;
    }
    public static void main(String[] args) {
        ListNode l1=new ListNode(1);
        l1.next=new ListNode(3);
        l1.next.next=new ListNode(4);

        ListNode l2=new ListNode(1);
        l2.next=new ListNode(4);
        l2.next.next=new ListNode(5);
        
        ListNode l3=new ListNode(2);
        l3.next=new ListNode(6);
        ListNode[] list={l1,l2,l3};
       mergeKLists(list);
    }
}