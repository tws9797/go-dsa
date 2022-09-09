package arrays

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = MergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}

func MergeTwoListsIterative(list1 *ListNode, list2 *ListNode) *ListNode {

	preHead := &ListNode{
		Val:  math.MinInt,
		Next: nil,
	}

	currentNode := preHead

	for currentNode != nil {

		if list1 == nil {
			currentNode.Next = list2
			break
		} else if list2 == nil {
			currentNode.Next = list1
			break
		} else if list1.Val > list2.Val {
			currentNode.Next = list2
			list2 = list2.Next
		} else {
			currentNode.Next = list1
			list1 = list1.Next
		}

		currentNode = currentNode.Next
	}

	return preHead.Next
}
