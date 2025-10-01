package march

import "fmt"

func reorderList(head *listNode) {
	if head.next == nil {
		return
	}
	s := head
	f := head.next

	for f.next != nil && f.next.next != nil {
		f = f.next.next
		s = s.next
	}

	if f.next != nil {
		s = s.next
	}
	f = s.next
	s.next = nil

	stack := []*listNode{}

	for f != nil {
		stack = append(stack, f)
		f = f.next
	}

	for l := len(stack) - 1; l >= 0; l-- {
		f = head.next
		head.next = stack[l]
		head.next.next = f
		head = f
	}
}

func Day23() {
	qs := []*listNode{
		arrIntoListNode([]int{1, 2, 3, 4}),
		arrIntoListNode([]int{1, 2, 3, 4, 5}),
	}

	for _, v := range qs {
		reorderList(v)
		fmt.Println(listNodeIntoArr(v))
	}
}
