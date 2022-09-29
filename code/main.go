package main

import (
	"QA/code/link"
	"fmt"
)

func main() {

	//nums := []int{1,2,6,3,4,5,6}
	head := &link.ListNode{Val: 1, Next: nil}
	head.Next = &link.ListNode{Val: 2, Next: nil}
	head.Next.Next = &link.ListNode{Val: 3, Next: nil}
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next

	}

}
