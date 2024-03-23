package march

import "fmt"

func isPalindrome(head *listNode) bool {
	if head.next == nil {
		return true
	}

	s := head
	f := head.next

	for f.next != nil && f.next.next != nil {
		s = s.next
		f = f.next.next
	}

	if f.next != nil {
		s = s.next
	}
	f = s.next
	s.next = nil
	f = reverseList(f)
	for f != nil {
		if f.val != head.val {
			return false
		}
		f = f.next
		head = head.next
	}
	return true
}

func Day22() {
	qs := []*listNode{
		arrIntoListNode([]int{1, 2, 3, 3, 2, 1}),
		arrIntoListNode([]int{1, 2}),
		arrIntoListNode([]int{1, 2, 1}),
	}

	for _, v := range qs {
		fmt.Println(isPalindrome(v))
	}
}
