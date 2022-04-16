/*

 * 234. Palindrome Linked List
 * Summit Count : 7 Times
 * Time used : 32 min

 * Solve by バンクー
 * Code with L#v

 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	result := isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2}})
	print(result)
}

func isPalindrome(current *ListNode) bool {
	var list = make([]int, 0)

	// Skip 1 index
	if current.Next == nil {
		return true
	}

	// Loop to List
	var next *ListNode = nil
	for {
		next = current.Next
		list = append(list, current.Val)
		if current.Next == nil {
			break
		}
		current = next
	}

	// Loop to Check
	for idx := 0; idx < len(list)-1; idx++ {
		if list[idx] != list[len(list)-1-idx] {
			return false
		}
	}

	return true
}
