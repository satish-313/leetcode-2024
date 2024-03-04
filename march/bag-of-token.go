package march

import (
	"fmt"
	"sort"
)

type day4 struct {
	tokens []int
	power  int
}

func bagOfTokenScore(tokens []int, power int) int {
	if len(tokens) == 0 {
		return 0
	}
	sort.Ints(tokens)
	l := 0
	r := len(tokens) - 1
	score := 0
	temp := 0
	if tokens[l] > power {
		return score
	}
	for l <= r {
		if power >= tokens[l] {
			power -= tokens[l]
			temp += 1
			l += 1
		} else if temp > 0 {
			power += tokens[r]
			temp -= 1
			r -= 1
		} else {
			break
		}
		if temp > score {
			score = temp
		}
	}
	return score
}

func Day4() {
	qa := []day4{
		{tokens: []int{100, 200, 300}, power: 50},
		{tokens: []int{200, 100}, power: 150},
		{tokens: []int{100, 200, 300, 400}, power: 200},
	}
	for _, v := range qa {
		fmt.Println(bagOfTokenScore(v.tokens, v.power))

	}
}
