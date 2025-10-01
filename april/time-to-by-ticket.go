package april

import (
	"fmt"
)

type day9 struct {
	tickets []int
	k       int
}

func timeRequiredToBuy(tickets []int, k int) int {
	t := 0
	v := tickets[k]

	for i := 0; i <= k; i++ {
		if v > tickets[i] {
			t += tickets[i]
		} else {
			t += v
		}
	}
	for i := k + 1; i < len(tickets); i++ {
		if v > tickets[i] {
			t += tickets[i]
		} else {
			t += v - 1
		}
	}
	return t
}

func Day9() {
	q := []day9{
		{[]int{2, 3, 2}, 2},
		{[]int{5, 1, 1, 1}, 0},
		{[]int{84, 49, 5, 24, 70, 77, 87, 8}, 3},
	}
	for _, v := range q {
		fmt.Println(timeRequiredToBuy(v.tickets, v.k))
	}
}
