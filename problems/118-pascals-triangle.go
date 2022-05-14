package main

/*

 * 118. Pascals-Triangle
 * Submit Count : 1 time
 * Time used :  20 minute

 * Solve by バンクー
 * Code with L#v

 */

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for row := 1; row <= numRows; row++ {
		record := make([]int, 0)
		for val := 0; val < row; val++ {
			if val != 0 && val != row-1 {
				record = append(record, result[row-2][val-1]+result[row-2][val])
			} else {
				record = append(record, 1)
			}
		}
		result = append(result, record)
	}
	return result
}

func main() {
	result := generate(5)
	fmt.Printf("%v", result)
}
