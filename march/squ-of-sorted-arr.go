package march

import (
	"fmt"
)

func sortedSquare(nums []int) []int {
	l := len(nums)

	if nums[l-1] < 0 {
		for i := 0; i < l/2; i++ {
			temp := nums[i] * nums[i]
			nums[i] = nums[l-1-i] * nums[l-1-i]
			nums[l-1-i] = temp
		}
		if nums[l/2] < 0 {
			nums[l/2] *= nums[l/2]
		}
		return nums
	}

	if nums[0] >= 0 {
		for i := 0; i < l; i++ {
			nums[i] *= nums[i]
		}
		return nums
	}

	for i := 1; i < l; i++ {
		if nums[i] >= 0 && nums[i-1] < 0 {
			nums[i-1] *= nums[i-1]
			return solve(i-1, i, nums)
		} else {
			nums[i-1] *= nums[i-1]
		}
	}

	return nums
}

func solve(l int, r int, nums []int) []int {
	res := make([]int, len(nums))
	idx := 0

	for l >= 0 && r < len(nums) {
		if nums[r]*nums[r] > nums[l] {
			res[idx] = nums[l]
			l -= 1
		} else {
			res[idx] = nums[r] * nums[r]
			r += 1
		}
		idx += 1
	}
	if r < len(nums) {
		for ; r < len(nums); r++ {
			res[r] = nums[r] * nums[r]
		}
	}
	if l >= 0 {
		for ; idx < len(nums); idx++ {
			res[idx] = nums[l]
			l -= 1
		}
	}
	return res
}

// func sortedSquaresbest(nums []int) []int {
// 	res := make([]int, 0, len(nums))
// 	left := 0
// 	right := len(nums) - 1

// 	for left <= right {
// 		if nums[left]*nums[left] > nums[right]*nums[right] {
// 			res = append(res, nums[left]*nums[left])
// 			left++
// 		} else {
// 			res = append(res, nums[right]*nums[right])
// 			right--
// 		}
// 	}
// 	reverseSlice(res)
// 	return res
// }

// func reverseSlice(nums []int) {
// 	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// }

func Day2() {
	qa := [][]int{{-12, -10, -9, -4, -1, 0, 3, 10}, {-7, 2, 3, 11, 13, 15, 16}}

	for _, val := range qa {
		fmt.Println(sortedSquare(val))
	}
}
