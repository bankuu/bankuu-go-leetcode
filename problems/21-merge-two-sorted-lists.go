package main

import (
	"bankuu-go-leetcode/problems/definition"
)

func mergeTwoLists(list1 definition.ListNode, list2 *definition.ListNode) *definition.ListNode {

	var result definition.ListNode

	var resultNext *definition.ListNode

	for {

		if list1.Next != nil && list2.Next != nil {
			if list1.Val <= list2.Val {

			} else {

			}
		} else if list1.Next != nil {

		} else if list2.Next != nil {

		} else {
			break
		}

	}

	return result

}
