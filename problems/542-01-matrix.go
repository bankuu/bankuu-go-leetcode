package main

/*

 * 542. 01 Matrix
 * Submit Count : 3 times
 * Time used : >3 Hours

 * Solve by バンクー
 * Code with L#v

 */

import "fmt"

func updateMatrix(mat [][]int) [][]int {
	//
	checker := 1

	for {
		changeCount := 0
		newMat := make([][]int, 0)
		for idxRow, row := range mat {
			newRow := make([]int, 0)

			for idxVal, _ := range row {
				isNearZero := false

				if mat[idxRow][idxVal] != checker {
					isNearZero = true
				}

				// check up
				if idxRow != 0 && mat[idxRow-1][idxVal] != checker {
					isNearZero = true
				}

				// check down
				if idxRow != len(mat)-1 && mat[idxRow+1][idxVal] != checker {
					isNearZero = true
				}

				// check left
				if idxVal != 0 && mat[idxRow][idxVal-1] != checker {
					isNearZero = true
				}

				// check right
				if idxVal != len(row)-1 && mat[idxRow][idxVal+1] != checker {
					isNearZero = true
				}

				if isNearZero == false {
					newRow = append(newRow, mat[idxRow][idxVal]+1)
					changeCount++
				} else {
					newRow = append(newRow, mat[idxRow][idxVal])
				}
			}
			newMat = append(newMat, newRow)
		}

		if changeCount == 0 {
			break
		}
		mat = newMat
		checker++
	}

	return mat
}

func main() {
	result := updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	})
	//result := updateMatrix([][]int{
	//	{0, 0, 1, 0, 1, 1, 1, 0, 1, 1},
	//	{1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
	//	{1, 1, 1, 1, 1, 0, 0, 0, 1, 1},
	//	{1, 0, 1, 0, 1, 1, 1, 0, 1, 1},
	//	{0, 0, 1, 1, 1, 0, 1, 1, 1, 1},
	//	{1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
	//	{1, 1, 1, 1, 0, 1, 0, 1, 0, 1},
	//	{0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
	//	{1, 1, 1, 0, 1, 1, 0, 1, 0, 1},
	//	{1, 0, 1, 1, 1, 0, 1, 1, 1, 0}})
	fmt.Printf("%v", result)
}
