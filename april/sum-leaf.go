package april

import (
	"fmt"
	"strconv"
)

func sumNumbers(root *treenode) int {
	sum := 0
	solve(root, "", &sum)
	return sum
}

func solve(r *treenode, s string, sum *int) {
	s += strconv.Itoa(r.Val)
	if r != nil && r.Left == nil && r.Right == nil {
		*sum += strInt(s)
		return
	}
	if r.Left != nil {
		solve(r.Left, s, sum)
	}
	if r.Right != nil {
		solve(r.Right, s, sum)
	}
}

func strInt(s string) int {
	ans, _ := strconv.Atoi(s)
	return ans
}

func Day15() {
	fmt.Println(sumNumbers(nil))
}
