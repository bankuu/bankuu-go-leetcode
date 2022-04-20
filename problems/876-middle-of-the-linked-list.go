/*

 * 876. Middle of the Linked List
 * Submit Count : 1 time
 * Time used : 41 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import (
	"bankuu-go-leetcode/problems/definition"
	"fmt"
)

func middleNode(head *definition.ListNode) *definition.ListNode {
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

	return definition.CreateListNode(intArr[len(intArr)/2:])
}

func main() {
	nodes := definition.CreateListNode([]int{1, 2, 3, 4, 5})
	result := middleNode(nodes)
	fmt.Printf("%v", result)
}
