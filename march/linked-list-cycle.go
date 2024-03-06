package march

import "fmt"

func hasCycle(head *listNode) bool {
	if head == nil {
		return false
	}
	l := head
	r := head.next

	if r == nil {
		return false
	}

	for r != nil && l != nil {
		if l == r {
			return true
		}
		l = l.next
		r = r.next.next
	}

	return false
}

func Day6() {
	arr := [][]int{{1, 2, 3, 4, 5}, {1}, {1, 2}}
	qa := []day3struct{
		{list: *arrIntoListNode(arr[0])},
		{list: *arrIntoListNode(arr[1])},
		{list: *arrIntoListNode(arr[2])},
	}
	fmt.Println(hasCycle(&qa[0].list))
}
