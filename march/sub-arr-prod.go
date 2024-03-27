package march

import (
	"fmt"
)

type day27 struct {
	arr  []int
	prod int
}

func numSubarrayProductLessThank(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	ans := 0
	prod := 1
	l := 0

	for r := 0; r < len(nums); r++ {
		prod *= nums[r]

		for prod >= k && l <= r {
			prod /= nums[l]
			l += 1
		}

		ans += r - l + 1
	}

	return ans
}

func Day27() {
	q := []day27{
		{arr: []int{10, 5, 2, 6}, prod: 100},
		{arr: []int{1, 2, 3}, prod: 0},
		{arr: []int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, prod: 19},
		{arr: []int{57, 44, 92, 28, 66, 60, 37, 33, 52, 38, 29, 76, 8, 75, 22}, prod: 18},
		{arr: []int{10, 2, 2, 5, 4, 4, 4, 3, 7, 7}, prod: 289},
	}

	for _, v := range q {
		fmt.Println(numSubarrayProductLessThank(v.arr, v.prod))
	}
}
