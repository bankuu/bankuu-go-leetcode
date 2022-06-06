package main

/*

 * 867. Transpose Matrix
 * Submit Count : 1 time
 * Time used : 40 m

 * Solve by バンクー
 * Code with L#v

 */

import "fmt"

func transpose(matrix [][]int) [][]int {
	var result = make([][]int, len(matrix[0]))
	for idxRows, rows := range matrix {
		for idx := range rows {
			result[idx] = append(result[idx], matrix[idxRows][idx])
		}
	}
	return result
}

func main() {
	result := transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Printf("%v", result)
}
