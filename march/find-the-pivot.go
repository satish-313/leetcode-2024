package march

import "fmt"

func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}
	sum := (n * (n + 1)) / 2
	pSum := 0

	for i := 1; i <= n; i++ {
		pSum += i
		if pSum == (sum - pSum + i) {
			return i
		}
	}

	return -1
}

func Day13() {
	q := []int{8, 1, 4}

	for _, val := range q {
		fmt.Println(pivotInteger(val))
	}
}
