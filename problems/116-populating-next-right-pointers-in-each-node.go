package main

/*

 * 116. Populating Next Right Pointers in Each Node
 * Submit Count : 1 time
 * Time used : 15 m

 * Solve by バンクー
 * Code with L#v

 */

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func nextConnect(listNode []*Node) {
	childNode := make([]*Node, 0)
	for idx := range listNode {
		if listNode[idx] != nil {
			// Replace Next point
			if idx < len(listNode)-1 {
				listNode[idx].Next = listNode[idx+1]
			}

			// Get Child to Replace
			childNode = append(childNode, listNode[idx].Left)
			childNode = append(childNode, listNode[idx].Right)
		}
	}

	if len(childNode) > 0 {
		nextConnect(childNode)
	}
}

func connect(root *Node) *Node {
	nextConnect([]*Node{root})
	return root
}

func main() {
	result := connect(&Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{Val: 3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	})

	fmt.Printf("%#v", result)
}
