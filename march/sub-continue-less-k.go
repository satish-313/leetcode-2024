package march

import "fmt"

func maxSubarrayLength(nums []int, k int) int {
	max := 0
	l := 0
	m := make(map[int]int)

	for r := 0; r < len(nums); r++ {
		m[nums[r]]++

		for m[nums[r]] > k && l <= r {
			m[nums[l]]--
			l += 1
		}

		if r-l+1 > max {
			max = r - l + 1
		}
	}

	return max
}

func Day28() {
	q := []day27{
		{arr: []int{1, 2, 3, 1, 2, 3, 12}, prod: 2},
	}

	for _, v := range q {
		fmt.Println(maxSubarrayLength(v.arr, v.prod))
	}
}
