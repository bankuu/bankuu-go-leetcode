/*

 * 19. Remove Nth Node From End of List
 * Submit Count : 1 time
 * Time used : 24 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import (
	"bankuu-go-leetcode/problems/definition"
	"fmt"
)

func removeNthFromEnd(head *definition.ListNode, n int) *definition.ListNode {
	intArr := make([]int, 0)
	node := head
	for {
		intArr = append(intArr, node.Val)
		if node.Next == nil {
			break
		} else {
			node = node.Next
		}
	}
	result := make([]int, 0)
	result = append(intArr[:len(intArr)-n])
	result = append(result, intArr[len(intArr)-n+1:]...)

	return definition.CreateListNode(result)
}

func main() {
	nodes := definition.CreateListNode([]int{1, 2})
	result := removeNthFromEnd(nodes, 1)
	fmt.Printf("%v", result)
}
