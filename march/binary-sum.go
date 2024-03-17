package march

import (
	"fmt"
)

type day14 struct {
	arr  []int
	goal int
}

func numSubarrayWithSum(nums []int, goal int) int {
	count := 0
	sum := 0
	l := 0
	r := 0

	for l <= r && r < len(nums) {
		sum += nums[r]

		for l <= r && sum > goal {
			sum -= nums[l]
			l += 1
		}

		if sum == goal {
			count += 1
		}
		if r == len(nums)-1 {
			for l <= r && r < len(nums) {
				sum -= nums[l]
				if sum == goal {
					count += 1
				}
				l += 1
			}
		}
		r += 1
	}

	return count
}

func Day14() {
	q := []day14{{arr: []int{1, 0, 1, 0, 1}, goal: 2}, {arr: []int{0, 0, 0, 0, 0}, goal: 0}}

	for _, val := range q {
		fmt.Println(numSubarrayWithSum(val.arr, val.goal))
	}
}
