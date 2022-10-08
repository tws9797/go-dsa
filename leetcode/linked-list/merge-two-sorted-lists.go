package linked_list

import "math"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	preNode := &ListNode{
		Val:  math.MinInt,
		Next: nil,
	}

	currentNode := preNode

	for currentNode != nil {

		if list1 == nil {
			currentNode.Next = list2
			break
		} else if list2 == nil {
			currentNode.Next = list1
			break
		} else if list1.Val < list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}

		currentNode = currentNode.Next
	}

	return preNode.Next
}
