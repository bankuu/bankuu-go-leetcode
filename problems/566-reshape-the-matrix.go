package main

/*

 * 557. Reshape the matrix
 * Submit Count : 3 times
 * Time used : 30 minutes

 * Solve by バンクー
 * Code with L#v

 */

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {

	result := make([][]int, 0)
	count := 0
	rCount := 0
	cCount := 0

	rowData := make([]int, 0)

	for _, rVal := range mat {
		for _, cVal := range rVal {

			if rCount >= r {
				count++
				continue
			}

			if cCount == 0 {
				rowData = make([]int, 0)
			}

			rowData = append(rowData, cVal)

			cCount++

			if cCount == c {
				cCount = 0
				result = append(result, rowData)
				rCount++
			}
			count++
		}
	}

	if count != r*c {
		return mat
	}

	return result
}

func main() {
	output := matrixReshape([][]int{{1, 2}}, 1, 1)
	fmt.Printf("%v", output)
}
