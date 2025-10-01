package april

import "fmt"

func makeGood(s string) string {
	stack := make([]rune, 0)

	for _, v := range s {
		if v > 64 && v < 91 {
			if len(stack) > 0 && stack[len(stack)-1] == v+32 {
				stack = stack[:len(stack)-1]
				continue
			}
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == v-32 {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, v)
	}
	return string(stack)
}

func Day5() {
	q := []string{
		"RLlr",
		"Pp",
		"leEeetcode",
		"abBAcC",
		"AZaz",
	}

	for _, v := range q {
		fmt.Println(v, makeGood(v))
	}
}
