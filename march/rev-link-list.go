package march

import "fmt"

func reverseList(head *listNode) *listNode {
	if head == nil {
		return head
	}
	var next *listNode
	var prev *listNode

	for head != nil {
		next = head.next
		head.next = prev
		prev = head
		head = next
	}

	return prev
}

func Day21() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(listNodeIntoArr(reverseList(arrIntoListNode(a))))
}
