package march

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	c1 := false

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			c1 = true
		}
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}

	if !c1 {
		return 1
	}

	for i := 0; i < len(nums); i++ {
		v := abs(nums[i])
		if v == len(nums) {
			nums[0] = -abs(nums[0])
		} else {
			nums[v] = -abs(nums[v])
		}
	}
	fmt.Println(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return len(nums)
	}

	return len(nums) + 1
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func Day26() {
	q := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
	}

	for _, v := range q {
		fmt.Println(firstMissingPositive(v))
	}
}
