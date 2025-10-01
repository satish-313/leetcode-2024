package april

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	result := make([]int, len(deck))
	stack := make([]int, 0)
	j := 0

	for i := 0; i < len(deck); i++ {
		if i%2 == 0 {
			result[i] = deck[j]
			j += 1
		} else {
			stack = append(stack, i)
		}
	}

	if len(deck)%2 != 0 && len(stack) > 0 {
		stack = append(stack, stack[0])
		stack = stack[1:]
	}

	for len(stack) > 0 {
		result[stack[0]] = deck[j]
		j += 1
		stack = stack[1:]
		if len(stack) > 0 {
			stack = append(stack, stack[0])
			stack = stack[1:]
		}
	}
	return result
}

func Day10() {
	q := [][]int{
		{17, 13, 11, 2, 5, 3, 7},
		{1, 1000},
		{1},
	}

	for _, v := range q {
		fmt.Println(deckRevealedIncreasing(v))
	}
}
