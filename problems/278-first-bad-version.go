/*

 * 278. First Bad Version
 * Submit Count : 2 times
 * Time used : 27 m

 * Solve by バンクー
 * Code with L#v

 */

package main

func isBadVersion(v int) bool {
	badVersion := 2
	return badVersion <= v
}

func firstBadVersion(n int) int {

	leftCursor := 1
	rightCursor := n

	for {
		center := (rightCursor + leftCursor) / 2

		if rightCursor-leftCursor <= 1 {
			if isBadVersion(leftCursor) {
				return leftCursor
			} else {
				return rightCursor
			}
		}

		isBad := isBadVersion(leftCursor)
		if isBad {
			if leftCursor == 1 || !isBadVersion(leftCursor-1) {
				return leftCursor
			}
			rightCursor = center
			leftCursor = center - rightCursor
		} else {
			leftCursor = center
		}
	}
}

func main() {
	result := firstBadVersion(10)
	print(result)
}
