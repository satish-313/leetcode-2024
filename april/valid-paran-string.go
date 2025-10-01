package april

import (
	"fmt"
)

func checkValidString(s string) bool {
	p1, p2 := 0, 0

	for _, v := range s {
		if v == '(' {
			p1++
			p2++
		} else if v == ')' {
			p1--
			p2--
		} else {
			p1++
			p2--
		}

		if p1 < 0 {
			return false
		}
		if p2 < 0 {
			p2 = 0
		}
	}
	return p2 == 0
}

func Day7() {
	q := []string{
		// "()",
		// "(*)",
		"(*))",
	}

	for _, v := range q {
		fmt.Println(checkValidString(v))
	}
}
