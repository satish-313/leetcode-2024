package april

import "fmt"

func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0)
	res := make([]byte, 0)
	c := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			c += 1
			stack = append(stack, i)
		} else if s[i] == ')' {
			if c > 0 {
				stack = stack[:len(stack)-1]
				c -= 1
			} else {
				stack = append(stack, i)
			}
		}
	}

	c = 0
	for i := 0; i < len(s); i++ {
		if c < len(stack) && stack[c] == i {
			c += 1
		} else {
			res = append(res, s[i])
		}
	}

	return string(res)
}

func Day6() {
	q := []string{
		"lee(t(c)o)de)",
		"a)b(c)d",
		"))((",
	}

	for _, v := range q {
		fmt.Println(minRemoveToMakeValid(v))
	}
}
