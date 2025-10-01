package march

import (
	"fmt"
)

func minLengthAfterOpe(s string) int {
	if len(s) == 0 {
		return 0
	}
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] == s[r] {
			for l < r && s[l] == s[l+1] {
				l += 1
			}
			for l < r && s[r] == s[r-1] {
				r -= 1
			}
		}

		if s[l] == s[r] {
			l += 1
			r -= 1
		} else {
			break
		}

	}

	if l > r {
		return 0
	}
	return r - l + 1
}

func Day5() {
	q := []string{"ca", "cabaabac", "aabccabba"}

	for _, v := range q {
		fmt.Println(minLengthAfterOpe(v))
	}
}
