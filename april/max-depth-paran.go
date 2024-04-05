package april

import (
	"fmt"
)

func maxDepths(s string) int {
	d := 0
	t := 0
	for _, v := range s {
		if v == '(' {
			t += 1
		} else if v == ')' {
			t -= 1
		}

		if t > d {
			d = t
		}
	}

	return d
}

func Day4() {
	q := []string{
		"(1+(2*3)+((8)/4))+1",
		"(1)+((2))+(((3)))",
	}

	for _, v := range q {
		fmt.Println(maxDepths(v))
	}
}
