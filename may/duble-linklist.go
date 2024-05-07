package may

import "fmt"

func doubleIt(head *Linknode) *Linknode {
	head = reverseLinkList(head)

	var prev *Linknode
	var next *Linknode
	carry := 0
	for head != nil || carry != 0 {
		if head == nil {
			h := Linknode{Val: carry, Next: prev}
			prev = &h
			break
		}
		val := (head.Val * 2) + carry
		fst := val % 10
		carry = val / 10
		head.Val = fst
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func Day7() {
	fmt.Println(doubleIt(nil))
}
