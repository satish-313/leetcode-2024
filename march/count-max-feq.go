package march

import "fmt"

func maxFrequencyElements(nums []int) int {
	m := make(map[int]int)
	mx := 0
	n := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		n = m[nums[i]]

		if mx < n {
			mx = n
			count = n
		} else if mx == n {
			count += n
		}
	}

	return count
}

func Day7() {
	qa := [][]int{{1, 2, 2, 3, 1, 4}, {1, 2, 3, 4, 5}}

	for _, v := range qa {
		fmt.Println(maxFrequencyElements(v))
	}
}
