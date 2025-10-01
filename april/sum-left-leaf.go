package april

import "fmt"

type treenode struct {
	Val   int
	Left  *treenode
	Right *treenode
}

func sumOfLeftLeaves(root *treenode) int {
	if root == nil {
		return 0
	}
	sum := 0
	queue := []*treenode{root}

	for len(queue) > 0 {
		n := len(queue)
		for _, v := range queue {
			if v.Left != nil {
				if v.Left.Left == nil && v.Left.Right == nil {
					sum += v.Left.Val
				}
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[n:]
	}

	return sum
}
func Day14() {
	fmt.Println(sumOfLeftLeaves(nil))
}
