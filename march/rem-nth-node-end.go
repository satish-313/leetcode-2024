package march

import "fmt"

type listNode struct {
	val  int
	next *listNode
}
type day3struct struct {
	list listNode
	nth  int
}

func removeNthFromEnd(head *listNode, n int) *listNode {
	if head == nil {
		return head
	}
	ans := head
	l := head
	r := head

	for i := 0; i < n; i++ {
		head = head.next
	}
	r = head

	if r == nil {
		return l.next
	}

	for r.next != nil {
		r = r.next
		l = l.next
	}
	if l.next.next == nil {
		l.next = nil
	} else {
		l.next = l.next.next
	}
	return ans
}

func arrIntoListNode(arr []int) *listNode {
	h := listNode{val: 0}
	c := &h

	for _, v := range arr {
		new := listNode{val: v}
		c.next = &new
		c = c.next
	}

	return h.next
}

func listNodeIntoArr(h *listNode) []int {
	if h == nil {
		return make([]int, 0)
	}

	ans := make([]int, 0)
	for h != nil {
		ans = append(ans, h.val)
		h = h.next
	}
	return ans
}

func Day3() {
	arr := [][]int{{1, 2, 3, 4, 5}, {1}, {1, 2}}
	qa := []day3struct{
		{list: *arrIntoListNode(arr[0]), nth: 2},
		{list: *arrIntoListNode(arr[1]), nth: 1},
		{list: *arrIntoListNode(arr[2]), nth: 1},
	}

	for _, v := range qa {
		fmt.Println(listNodeIntoArr(removeNthFromEnd(&v.list, v.nth)))
	}
}
