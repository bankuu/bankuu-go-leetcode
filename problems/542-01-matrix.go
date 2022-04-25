package main

import "fmt"

type locate struct {
	x int
	y int
}

func checkUpdateMatrix(mat [][]int, checkList map[locate]*int, checkUp int) [][]int {
	newValList := make(map[locate]*int)

	for key, cal := range checkList {
		//// Check Left
		//if key.x != 0 && (mat)[key.y][key.x-1] < checkUp {
		//	continue
		//}
		//// Check Top
		//if key.y != 0 && (mat)[key.y-1][key.x] < checkUp {
		//	continue
		//}
		//
		//// Check Right
		//if key.x != len((mat)[key.x])+1 && (mat)[key.y][key.x+1] < checkUp {
		//	continue
		//}
		//
		//// Check Down
		//if key.y != len((mat)[key.y])+1 && (mat)[key.y+1][key.x] < checkUp {
		//	continue
		//}
		//
		//newVal := (*checkList[key]) + 1
		//checkList[key] = &newVal
		//mat[key.y][key.x] = newVal
		//newValList[locate{x: key.x, y: key.y}] = &newVal
	}

	if len(newValList) > 0 {
		mat = checkUpdateMatrix(mat, newValList, checkUp+1)
	}

	return mat

}

func updateMatrix(mat [][]int) [][]int {
	checkList := make(map[locate]*int)
	for y, row := range mat {
		for x, val := range row {
			checkList[locate{x: x, y: y}] = &val
		}
	}
	return checkUpdateMatrix(mat, checkList, 1)

}

func main() {
	result := updateMatrix([][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	})

	fmt.Printf("%v", result)
}
