package march

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	a := 0
	b := 0
	ans := make([]int, 0)
	for a < len(nums1) && b < len(nums2) {
		if nums1[a] == nums2[b] {
			if len(ans) == 0 {
				ans = append(ans, nums1[a])
			} else if ans[len(ans)-1] != nums1[a] {
				ans = append(ans, nums1[a])
			}
			a += 1
			b += 1
		} else if nums1[a] > nums2[b] {
			b += 1
		} else {
			a += 1
		}
	}

	return ans
}

func Day10() {
	fmt.Println(intersection([]int{1, 2, 2, 4, 2, 3}, []int{3, 1, 4, 2, 5, 6}))
}
