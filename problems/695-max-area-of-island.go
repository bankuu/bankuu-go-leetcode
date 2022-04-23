package main

/*

 * 695. Max Area of island
 * Submit Count : 1 time
 * Time used : 18 m

 * Solve by バンクー
 * Code with L#v

 */

func countFloodFill(image [][]int, sr int, sc int, newColor int, count *int) {
	oldColor := image[sr][sc]

	// Left
	if sc-1 >= 0 && image[sr][sc-1] == oldColor && image[sr][sc-1] != newColor {
		countFloodFill(image, sr, sc-1, newColor, count)
	}
	if image[sr][sc] != newColor {
		image[sr][sc] = newColor
		*count = *count + 1
	}

	if sc+1 < len(image[sr]) && image[sr][sc+1] == oldColor && image[sr][sc+1] != newColor {
		countFloodFill(image, sr, sc+1, newColor, count)
	}
	if image[sr][sc] != newColor {
		image[sr][sc] = newColor
		*count = *count + 1
	}

	//Top
	if sr-1 >= 0 && image[sr-1][sc] == oldColor && image[sr-1][sc] != newColor {
		countFloodFill(image, sr-1, sc, newColor, count)
	}
	if image[sr][sc] != newColor {
		image[sr][sc] = newColor
		*count = *count + 1
	}

	//
	if sr+1 < len(image) && image[sr+1][sc] == oldColor && image[sr+1][sc] != newColor {
		countFloodFill(image, sr+1, sc, newColor, count)
	}
	if image[sr][sc] != newColor {
		image[sr][sc] = newColor
		*count = *count + 1
	}

	return
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for x, _ := range grid {
		for y, _ := range grid[x] {
			if grid[x][y] == 1 {
				count := 0
				countFloodFill(grid, x, y, 2, &count)
				if count > maxArea {
					maxArea = count
				}
			}
		}
	}
	return maxArea
}

func main() {
	max := maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	})
	print(max)
}
