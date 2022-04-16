/*

 * 2. Add Two Number
 * Summit Count : 3 times
 * Time used : 52 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import (
	"fmt"
	"go/token"
	"go/types"
	"strconv"
	"strings"
)

func main() {
	thing := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	matter := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	addTwoNumbers(&thing, &matter)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := listNodeToString(l1)
	s2 := listNodeToString(l2)
	fs := token.NewFileSet()
	tv, _ := types.Eval(fs, nil, token.NoPos, fmt.Sprintf("%s + %s", s1, s2))
	return stringToListNode(tv.Value.String())
}

func listNodeToString(l *ListNode) string {
	str := ""
	str = fmt.Sprintf("%s%d", str, l.Val)
	for l.Next != nil {
		str = fmt.Sprintf("%d%s", l.Next.Val, str)
		l = l.Next
	}
	return str
}

func stringToListNode(str string) *ListNode {
	var result *ListNode
	for _, char := range strings.Split(str, "") {
		i, _ := strconv.Atoi(char)
		if result != nil {
			result = &ListNode{Val: i, Next: result}
		} else {
			result = &ListNode{Val: i}
		}
	}
	return result
}
