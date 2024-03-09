package march

import "fmt"

func getCommon(nums1 []int, nums2 []int) int {
	a := 0
	b := 0

	for a < len(nums1) && b < len(nums2) {
		if nums1[a] == nums2[b] {
			return nums1[a]
		} else if nums1[a] > nums2[b] {
			b += 1
		} else {
			a += 1
		}
	}

	return -1
}

func Day9() {
	fmt.Println(getCommon([]int{1, 2, 3, 4, 5}, []int{2, 3, 5, 9}))
}
