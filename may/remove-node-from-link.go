package may

import (
	"fmt"
)

func removeNodes(head *Linknode) *Linknode {
	arr := make([]int, 0)
	cur := head
	res := Linknode{Val: 0}
	ans := &res

	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	max := -1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < max {
			continue
		}
		max = arr[i]
		new := Linknode{Val: arr[i], Next: nil}
		ans.Next = &new
		ans = ans.Next
	}

	return reverseLinkList(res.Next)
}
func Day6() {
	q := []*Linknode{
		arrToLinkNode([]int{5, 2, 13, 3, 8}),
		// *arrToLinkNode([]int{1,1,1,1}),
	}

	for _, v := range q {
		fmt.Println(linkToArr(removeNodes(v)))
	}
}
