package main

import (
	"bankuu-go-leetcode/problems/definition"
	"fmt"
	"math"
)

func mergeTrees(root1 *definition.TreeNode, root2 *definition.TreeNode) *definition.TreeNode {
	var resultAry []*int

	root1Ary := treeNodeToArray(root1)
	root2Ary := treeNodeToArray(root2)

	if len(root1Ary) > len(root2Ary) {
		resultAry = make([]*int, len(root1Ary))
	} else {
		resultAry = make([]*int, len(root2Ary))
	}

	for idx := range resultAry {
		if (len(root1Ary) > idx && root1Ary[idx] != nil) || (len(root2Ary) > idx && root2Ary[idx] != nil) {
			val := 0
			if len(root1Ary) > idx && root1Ary[idx] != nil {
				val = val + *root1Ary[idx]
			}
			if len(root2Ary) > idx && root2Ary[idx] != nil {
				val = val + *root2Ary[idx]
			}
			resultAry[idx] = &val
		}
	}

	return ArrayToTreeNode(resultAry)
}

func treeNodeToArray(root *definition.TreeNode) []*int {
	result := make([]*int, 0)
	stackNode := []*definition.TreeNode{root}

	for {
		// check null layer
		nullCount := 0
		childNode := make([]*definition.TreeNode, 0)

		//for to write value
		for _, val := range stackNode {
			if val != nil {
				result = append(result, &val.Val)
			} else {
				result = append(result, nil)
				continue
			}

			// check left
			if val.Left != nil {
				childNode = append(childNode, val.Left)
			} else {
				childNode = append(childNode, nil)
				nullCount++
			}

			// check right
			if val.Right != nil {
				childNode = append(childNode, val.Right)
			} else {
				childNode = append(childNode, nil)
				nullCount++
			}
		}

		if nullCount >= len(childNode) {
			return result
		} else {
			stackNode = childNode
		}
	}

}

func ArrayToTreeNode(val []*int) *definition.TreeNode {
	mapper := make(map[int][]*definition.TreeNode)

	row := 0
	innerSum := 0
	for {
		create := make([]*definition.TreeNode, 0)
		inner := int(math.Pow(2, float64(row)))
		for idx := innerSum; idx <= innerSum+inner-1; idx++ {
			if val[idx] != nil {
				create = append(create, &definition.TreeNode{Val: *val[idx]})
			} else {
				create = append(create, nil)
			}
		}

		innerSum = innerSum + inner

		mapper[row] = create
		if innerSum >= len(val) {
			break
		}
		row++
	}

	for key := range mapper {
		if key >= len(mapper)-1 {
			break
		}

		for idx, _ := range mapper[key+1] {
			parentPos := idx / 2
			if idx%2 == 0 {
				mapper[key][parentPos].Left = mapper[key+1][idx]
			} else {
				mapper[key][parentPos].Right = mapper[key+1][idx]
			}
		}

	}
	return mapper[0][0]
}

func main() {
	root1 := &definition.TreeNode{
		Val: 1,
		Left: &definition.TreeNode{
			Val: 2,
			Left: &definition.TreeNode{
				Val: 3,
			},
		},
	}

	root2 := &definition.TreeNode{
		Val: 1,
		Left: &definition.TreeNode{
			Val:  0,
			Left: nil,
			Right: &definition.TreeNode{
				Val: 0,
				Right: &definition.TreeNode{
					Val: 3,
				},
			},
		},
		Right: &definition.TreeNode{
			Val: 2,
		},
	}

	result := mergeTrees(root1, root2)
	fmt.Printf("%#v", result)
}
