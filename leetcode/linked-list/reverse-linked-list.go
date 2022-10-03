package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/reverse-linked-list/

func ReverseList(head *ListNode) *ListNode {

	var prev *ListNode
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}

	return prev
}

// 1->2->3->4->5

func ReverseListRecursively(head *ListNode) *ListNode {

	// head == nil for []
	// head.Next for head.Next.Next
	if head == nil || head.Next == nil {
		return head
	}
	r := ReverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return r
}
