package addtwonums

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(l []int) *ListNode {
	ln := &ListNode{Val: l[0]}
	start := ln
	for i := 1; i < len(l); i++ {
		ln.Next = &ListNode{Val: l[i]}
		ln = ln.Next
	}
	return start
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var one, two float64
	mul := float64(1)

	for ; l1 != nil || l2 != nil; mul = mul * 10 {
		if l1 != nil {
			one = one + (float64(l1.Val) * mul)
			l1 = l1.Next
		}
		if l2 != nil {
			two = two + (float64(l2.Val) * mul)
			l2 = l2.Next
		}
	}
	res := int(one + two)
	var start, l3 *ListNode

	if res == 0 {
		return &ListNode{Val: 0, Next: nil}
	}

	for res > 0 {
		if l3 != nil {
			l3.Next = &ListNode{}
			l3 = l3.Next
		} else {
			l3 = &ListNode{}
			start = l3
		}
		val := res % 10
		l3.Val = val
		res = (res - val) / 10
	}
	return start
}

func addTwoNumbersNoMul(l1 *ListNode, l2 *ListNode) *ListNode {
	var one, two, rem int
	var l3, start *ListNode

	for l1 != nil || l2 != nil || rem != 0 {
		if l3 != nil {
			l3.Next = &ListNode{}
			l3 = l3.Next
		} else {
			l3 = &ListNode{}
			start = l3
		}

		one, two = 0, 0

		if l1 != nil {
			one = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			two = l2.Val
			l2 = l2.Next
		}

		sum := one + two
		val := (sum % 10) + (rem % 10)
		l3.Val = val % 10
		rem = ((val - (val % 10)) + (sum - (sum % 10)) + (rem - (rem % 10))) / 10
	}
	return start
}
