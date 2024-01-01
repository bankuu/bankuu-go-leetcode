package solutions

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	maxCount := 0

	var recList []string

	if len(s) == 1 {
		return 1
	}

	for _, bary := range s {
		// check char
		var newRecList []string

		for recIdx, rec := range recList {
			if strings.Contains(recList[recIdx], string(bary)) {
				if len(rec) > maxCount {
					maxCount = len(rec)
				}
			} else {
				newRecList = append(newRecList, rec)
			}
		}

		// replace rec
		recList = newRecList

		// add to rec
		for recIdx := range recList {
			if len(recList[recIdx]) > 0 {
				recList[recIdx] = fmt.Sprintf("%s%s", recList[recIdx], string(bary))
			}
		}

		// add new rec
		recList = append(recList, string(bary))
	}

	// loop find after
	for idx := range recList {
		if len(recList[idx]) > maxCount {
			maxCount = len(recList[idx])
		}
	}

	return maxCount
}
