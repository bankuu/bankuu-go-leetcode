/*

 * 733. Flood Fill
 * Submit Count : 1 time
 * Time used : 50 m

 * Solve by バンクー
 * Code with L#v

 */

package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	oldColor := image[sr][sc]

	// Left
	if sc-1 >= 0 && image[sr][sc-1] == oldColor && image[sr][sc-1] != newColor {
		floodFill(image, sr, sc-1, newColor)
	}
	image[sr][sc] = newColor

	if sc+1 < len(image[sr]) && image[sr][sc+1] == oldColor && image[sr][sc+1] != newColor {
		floodFill(image, sr, sc+1, newColor)
	}
	image[sr][sc] = newColor

	//Top
	if sr-1 >= 0 && image[sr-1][sc] == oldColor && image[sr-1][sc] != newColor {
		image = floodFill(image, sr-1, sc, newColor)
	}
	image[sr][sc] = newColor

	//
	if sr+1 < len(image) && image[sr+1][sc] == oldColor && image[sr+1][sc] != newColor {
		image = floodFill(image, sr+1, sc, newColor)
	}

	return image
}

func main() {
	result := floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2)
	fmt.Printf("%v", result)
}
