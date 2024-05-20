package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	pre := &ListNode{}
	cur := pre
	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return pre.Next
}

func main() {
	var l1 = ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var l2 = ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	addTwoNumbers(&l1, &l2)
}
