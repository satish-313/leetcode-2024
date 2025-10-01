package april

import "fmt"

func lengthOfLastWord(s string) int {
	l := 0
	e := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			e = true
			l += 1
		}
		if s[i] == ' ' && e {
			return l
		}
	}

	return l
}

func Day1() {
	qs := []string{
		"Hello world",
		"  fly me  to  the moon  ",
		"luffy is still joyboy",
	}

	for _, v := range qs {
		fmt.Println(lengthOfLastWord(v))
	}
}
