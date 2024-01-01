package solutions

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// ListNode singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Str := ""
	l2Str := ""
	for l1 != nil {
		l1Str = fmt.Sprintf("%d%s", l1.Val, l1Str)
		l1 = l1.Next
	}
	for l2 != nil {
		l2Str = fmt.Sprintf("%d%s", l2.Val, l2Str)
		l2 = l2.Next
	}
	l1Int, _ := new(big.Int).SetString(l1Str, 10)
	l2Int, _ := new(big.Int).SetString(l2Str, 10)
	result := l1Int.Add(l1Int, l2Int)

	var rList *ListNode
	for _, rStr := range strings.Split(fmt.Sprintf("%s", result.String()), "") {
		rInt, _ := strconv.Atoi(rStr)
		rList = &ListNode{Val: rInt, Next: rList}
	}
	return rList
}
