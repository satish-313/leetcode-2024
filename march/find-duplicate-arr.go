package march

import "fmt"

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n < 0 {
			n *= -1
		}
		if nums[n-1] < 0 {
			res = append(res, n)
		}
		nums[n-1] *= -1
	}

	return res
}

func Day25() {
	q := [][]int{
		{4, 3, 2, 7, 8, 2, 3, 1},
		{1, 1, 2},
		{1},
	}

	for _, v := range q {
		fmt.Println(findDuplicates(v))
	}
}
