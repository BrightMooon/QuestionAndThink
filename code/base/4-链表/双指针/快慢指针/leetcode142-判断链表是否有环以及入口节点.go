package link
//题解 https://leetcode.cn/problems/linked-list-cycle-ii/solution/yi-yan-jiu-kan-dong-de-ti-jie-shuang-zhi-4sag/
//1 hash解法
//func detectCycle(head *ListNode) *ListNode {
//     if head==nil{
//         return nil
//     }
//     if head.Next==nil{
//         return nil
//     }
//     nodeMap:=make(map[*ListNode]int)
//     for head!=nil{
//         if _,ok:=nodeMap[head];ok{
//             return head
//         }
//         nodeMap[head]=1
//         head=head.Next
//     }
//     return nil
// }
//2 双指针，^相遇后，新节点从head开始移动，low从相遇节点开始移动，再次相遇时，就是环的入口节点
func detectCycle(head *ListNode) *ListNode {
	fast,slow:=head,head
	for fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			begin:=head
			for begin!=slow{
				begin=begin.Next
				slow=slow.Next
			}
			return slow
		}
	}
	return nil
}
