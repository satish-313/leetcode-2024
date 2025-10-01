package april

import "fmt"

func addOneRow(root *treenode, val int, depth int) *treenode {
	return solve16(root, val, depth, true)
}

func solve16(root *treenode, v int, d int, isLeft bool) *treenode {
	if d == 1 {
		r := &treenode{Val: v}
		if isLeft {
			r.Left = root
		} else {
			r.Right = root
		}
		return r
	}
	if root == nil {
		return root
	}
	d -= 1
	root.Left = solve16(root.Left, v, d, true)
	root.Right = solve16(root.Right, v, d, false)

	return root
}

func Day16() {
	fmt.Println(addOneRow(nil, 0, 0))
}
