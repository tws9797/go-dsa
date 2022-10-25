package linked_list

// https://leetcode.com/problems/linked-list-cycle/

func HasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	m := map[*ListNode]bool{}

	for head.Next != nil {

		if _, ok := m[head]; ok {
			return true
		}

		m[head] = true
		head = head.Next
	}

	return false
}

func HasCycleConstantSpace(head *ListNode) bool {

	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {

		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next

	}

	return true
}
