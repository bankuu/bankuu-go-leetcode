package solutions

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	longestText := string(s[0])
	for pos := (len(s) + 1) / 2; pos < len(s); pos++ {
		lpos := len(s) - pos
		rpos := pos

		fmt.Printf("lpos : %d, rpos: %d\n", lpos, rpos)

		find := func(pos int) {
			iasq := true
			ibsq := true
			shift := 1
			for {
				lcur := pos - shift
				arcur := pos + shift
				brcur := pos + shift + 1

				if lcur < 0 || (!iasq && !ibsq) {
					// get longest find
					//fmt.Printf("longest %d\n", pos)
					break
				}

				if arcur > len(s) || s[lcur] != s[arcur-1] {
					iasq = false
				} else if iasq {
					if len(longestText) < len(s[lcur:arcur]) {
						longestText = s[lcur:arcur]
					}
					fmt.Printf("lcur:%d [%s] arcur:%d [%s] res:%s \n", lcur, string(s[lcur]), arcur, string(s[arcur-1]), s[lcur:arcur])
				}

				if brcur > len(s) || s[lcur] != s[brcur-1] {
					ibsq = false
				} else if ibsq {
					if len(longestText) < len(s[lcur:brcur]) {
						longestText = s[lcur:brcur]
					}
					fmt.Printf("lcur:%d [%s] brcur:%d [%s] res:%s \n", lcur, string(s[lcur]), arcur, string(s[brcur-1]), s[lcur:brcur])
				}
				shift++
			}
		}
		find(rpos)
		if rpos != lpos {
			find(lpos)
		}

	}
	return longestText
}
