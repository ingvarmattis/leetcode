package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current := &ListNode{}
	root := current
	temp := 0
	for {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += temp

		current.Val = sum % 10
		temp = sum / 10

		if l1 == nil && l2 == nil && temp == 0 {
			break
		}

		current.Next = &ListNode{}
		current = current.Next
	}

	return root
}
