package march

import (
	"fmt"
	"strings"
)

func customSortString(order, s string) string {
	sArr := [26]int{}
	oArr := [26]int{}
	var ans strings.Builder

	for i := 0; i < len(s); i++ {
		sArr[s[i]-97]++
	}

	for i := 0; i < len(order); i++ {
		oArr[order[i]-97]++
		if sArr[order[i]-97] != 0 {
			for j := 0; j < sArr[order[i]-97]; j++ {
				fmt.Fprintf(&ans, "%s", string(order[i]))
			}
		}
	}
	for i := 0; i < len(sArr); i++ {
		if sArr[i] != 0 && oArr[i] == 0 {
			for j := 0; j < sArr[i]; j++ {
				fmt.Fprintf(&ans, "%s", string(i+97))
			}
		}
	}

	return ans.String()
}

func Day11() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("bcafg", "abcd"))
}
