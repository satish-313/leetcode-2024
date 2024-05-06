package may

type Linknode struct {
	Val  int
	Next *Linknode
}

func arrToLinkNode(arr []int) *Linknode {
	h := Linknode{Val: 0}
	c := &h

	for _, v := range arr {
		new := Linknode{Val: v}
		c.Next = &new
		c = c.Next
	}

	return h.Next
}

func linkToArr(head *Linknode) []int {
	arr := make([]int, 0)

	if head == nil {
		return nil
	}

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return arr
}

func reverseLinkList(head *Linknode) *Linknode {
	var prev *Linknode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
