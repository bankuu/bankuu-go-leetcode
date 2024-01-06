package solutions

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	isMinus := false
	if x <= 0 {
		x = x * -1
		isMinus = true
	}
	str := ""
	for _, char := range strconv.Itoa(x) {
		str = fmt.Sprintf("%s%s", string(char), str)
	}
	if isMinus {
		str = fmt.Sprintf("-%s", str)
	}
	result, _ := strconv.Atoi(str)

	if result > 2147483647 || result < -2147483647 {
		result = 0
	}

	return result
}
