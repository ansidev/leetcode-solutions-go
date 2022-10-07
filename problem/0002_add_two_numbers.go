package problem

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummyNode := &ListNode{}
	result := dummyNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := valueOrZero(l1) + valueOrZero(l2) + carry
		carry = sum / 10

		// this will modify to dummyNode value because result is referencing to dummyNode
		result.Next = &ListNode{
			Val: sum % 10,
		}

		result = result.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		result.Next = &ListNode{Val: carry}
	}

	return dummyNode.Next
}

func valueOrZero(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}
