package solutions

import "fmt"

func convert(s string, numRows int) string {
	recString := make([]string, numRows)
	pos := 0
	isReverse := false

	// skip if numRows == 1
	if numRows == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {

		// record data
		recString[pos] = fmt.Sprintf("%s%s", recString[pos], string(s[i]))

		// -- control pos
		if isReverse {
			pos--
		} else {
			pos++
		}

		if pos == numRows-1 {
			isReverse = true
		} else if pos == 0 {
			isReverse = false
		}
		// --
	}

	// re-index to result
	result := ""
	for _, rec := range recString {
		result = fmt.Sprintf("%s%s", result, rec)
	}

	return result
}
