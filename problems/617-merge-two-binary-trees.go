package main

/*

 * 617. Merge Two Binary Trees
 * Submit Count : 2 times
 * Time used : > 3 hours

 * Solve by バンクー
 * Code with L#v

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var node *TreeNode

	// Check Value
	if root1 != nil || root2 != nil {

		// Replace Value
		node = &TreeNode{Val: 0}
		if root1 != nil {
			node.Val = node.Val + root1.Val
		}
		if root2 != nil {
			node.Val = node.Val + root2.Val
		}
		// --

		// Turn Left

		var root1Left *TreeNode
		var root2Left *TreeNode

		if root1 != nil && root1.Left != nil {
			root1Left = root1.Left
		}

		if root2 != nil && root2.Left != nil {
			root2Left = root2.Left
		}

		node.Left = mergeTrees(root1Left, root2Left)

		// Turn Right
		var root1Right *TreeNode
		var root2Right *TreeNode

		if root1 != nil && root1.Right != nil {
			root1Right = root1.Right
		}

		if root2 != nil && root2.Right != nil {
			root2Right = root2.Right
		}

		node.Right = mergeTrees(root1Right, root2Right)

	}

	return node
}

func main() {
	result := mergeTrees(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}, &TreeNode{Val: 2, Left: &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 4},
	}, Right: &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 7},
	}})
	print(result)
}
