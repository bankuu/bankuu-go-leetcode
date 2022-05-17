package main

/*

 * 994. Rotting Oranges
 * Submit Count : 1 time
 * Time used : 20 m

 * Solve by バンクー
 * Code with L#v

 */

import (
	"fmt"
	"strconv"
)

func orangesRotting(grid [][]int) int {
	route := 0

	for {
		countOne := 0

		var idxList = make(map[string]interface{})

		for rowIdx, rowVal := range grid {
			for idx, val := range rowVal {
				if val != 2 {
					if val == 1 {
						countOne++
					}
					continue
				} else {
					//check round is one

					// left
					if idx > 0 && grid[rowIdx][idx-1] == 1 {
						idxList[fmt.Sprintf("%d-%d", rowIdx, idx-1)] = nil
					}

					// right
					if idx < len(rowVal)-1 && grid[rowIdx][idx+1] == 1 {
						idxList[fmt.Sprintf("%d-%d", rowIdx, idx+1)] = nil
					}

					// top
					if rowIdx > 0 && grid[rowIdx-1][idx] == 1 {
						idxList[fmt.Sprintf("%d-%d", rowIdx-1, idx)] = nil
					}

					// bottom
					if rowIdx < len(grid)-1 && grid[rowIdx+1][idx] == 1 {
						idxList[fmt.Sprintf("%d-%d", rowIdx+1, idx)] = nil
					}
				}
			}
		}

		if len(idxList) == 0 {
			if countOne > 0 {
				return -1
			}
			break
		} else {
			for idx, _ := range idxList {
				rowIdx, _ := strconv.Atoi(string(idx[0]))
				idx, _ := strconv.Atoi(string(idx[2]))
				grid[rowIdx][idx] = 2
			}
			route++
		}
	}

	return route
}

func main() {
	result := orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})
	print(result)
}
