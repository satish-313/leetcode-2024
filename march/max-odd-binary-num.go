package march

import (
	"fmt"
	"strings"
)

func maximumOddBinaryNumber(s string) string {
	ones := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones++
		}
	}

	arr := make([]string, len(s))

	left := 0
	arr[len(s)-1] = "1"
	ones -= 1

	for ones > 0 {
		arr[left] = "1"
		ones -= 1
		left += 1
	}
	for i := 0; i < len(s); i++ {
		if arr[i] != "1" {
			arr[i] = "0"
		}
	}
	return strings.Join(arr, "")
}

func Day1() {
	qs := []string{"010", "0101"}
	for i := 0; i < len(qs); i++ {
		fmt.Println(maximumOddBinaryNumber(qs[i]))
	}
}
