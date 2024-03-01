package feb

import (
	"fmt"
)

func findJudge(n int, trust [][]int) int {
	l, r := make([]int, n+1), make([]int, n+1)

	for _, t := range trust {
		l[t[0]]++
		r[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if l[i] == 0 && r[i] == n-1 {
			return i
		}
	}
	return -1
}

func Day22() {
	n := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	fmt.Println(findJudge(n, trust))
}
