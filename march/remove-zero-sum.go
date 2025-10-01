package march

import (
	"fmt"
)

func removeZeroSumSublists(head *listNode) *listNode {
	front := listNode{0, head}
	cur := &front
	prefixSum := 0
	m := make(map[int]*listNode)

	for cur != nil {
		prefixSum += cur.val

		v, ok := m[prefixSum]
		if ok {
			prev := v
			cur = prev.next

			p := prefixSum + cur.val
			for p != prefixSum {
				delete(m, p)
				cur = cur.next
				p += cur.val
			}
			prev.next = cur.next
		} else {
			m[prefixSum] = cur
		}
		cur = cur.next
	}

	return front.next
}

func Day12() {
	q := [][]int{{1, 2, -3, 3, 1}, {1, 2, 3, -3, 4}, {1, 2, 3, -3, -2}, {5, -3, -4, 1, 6, -2, -5}}
	qs := []listNode{
		*arrIntoListNode(q[0]),
		*arrIntoListNode(q[1]),
		*arrIntoListNode(q[2]),
		*arrIntoListNode(q[3]),
	}
	for _, v := range qs {
		fmt.Println(listNodeIntoArr(removeZeroSumSublists(&v)))
	}
}
