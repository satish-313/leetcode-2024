package april

import (
	"fmt"
	"strings"
)

func smallestFromLeaf(root *treenode) string {
	sm := ""
	solve17(root, "", &sm)
	return sm
}

func solve17(r *treenode, s string, sm *string) {
	s += fmt.Sprintf("%c", r.Val+97)
	if r.Left == nil && r.Right == nil {
		rs := reverse(&s)

		if len(*sm) == 0 {
			*sm = rs
		}
		if rs < *sm {
			*sm = rs
		}
		return
	}

	if r.Left != nil {
		solve17(r.Left, s, sm)
	}
	if r.Right != nil {
		solve17(r.Right, s, sm)
	}
}

func reverse(s *string) string {
	var b strings.Builder

	for i := len(*s) - 1; i >= 0; i-- {
		fmt.Fprintf(&b, "%c", (*s)[i])
	}

	return b.String()
}

func Day17() {
	fmt.Println(smallestFromLeaf(nil))
}
