package solutions

// ListNode singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	for l1.Next != nil {
		l1 = l1.Next
	}
	return nil
}
