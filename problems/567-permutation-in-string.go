package main

/*

 * 567. Permutation in String
 * Submit Count : 5 times
 * Time used : 42 m

 * Solve by バンクー
 * Code with L#v

 */

import "strings"

func checkInclusion(s1 string, s2 string) bool {
	lPointer := 0
	rPointer := len(s1)

	for rPointer <= len(s2) {
		strCompare := s1
		strPointer := s2[lPointer:rPointer]
		for pos := len(strPointer); pos > 0; pos-- {
			var idx = pos - 1
			if strings.Contains(strCompare, string(strPointer[idx])) {
				strCompare = strings.Replace(strCompare, string(strPointer[idx]), "", 1)
			} else {
				lPointer = lPointer + len(strCompare) - 1
				rPointer = rPointer + len(strCompare) - 1
				break
			}
		}

		if len(strCompare) == 0 {
			return true
		}
		lPointer = lPointer + 1
		rPointer = rPointer + 1
	}
	return false
}

func main() {
	output := checkInclusion("abc", "bbbca")
	print(output)
}
