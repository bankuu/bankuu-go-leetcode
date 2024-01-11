package solutions

import (
	"fmt"
	"math/big"
	"strings"
)

func myAtoi(s string) int {
	checkStr := "0123456789+-"

	record := ""
	isRecording := true

	for _, char := range s {
		if len(record) > 0 && (string(char) == "-" || string(char) == "+") {
			break
		}
		for checkIdx, check := range checkStr {
			if char == check {
				record = fmt.Sprintf("%s%s", record, string(check))
				isRecording = true
				break
			} else if checkIdx == len(checkStr)-1 {
				isRecording = false
				break
			}
		}

		if (len(record) > 0 || string(char) != " ") && isRecording == false {
			break
		}
	}

	if result, ok := new(big.Int).SetString(strings.TrimSpace(record), 10); ok == true {
		if result.Cmp(big.NewInt(2147483647)) > 0 {
			return 2147483647
		} else if result.Cmp(big.NewInt(-2147483648)) < 0 {
			return -2147483648
		} else {
			return int(result.Int64())
		}
	} else {
		return 0
	}
}
