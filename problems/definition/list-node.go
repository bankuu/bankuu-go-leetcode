package definition

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(intAry []int) (head *ListNode) {
	var node *ListNode
	for _, val := range intAry {
		tail := ListNode{Val: val}
		if head == nil {
			head = &tail
			node = head
		} else {
			node.Next = &tail
		}
		node = &tail
	}
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
