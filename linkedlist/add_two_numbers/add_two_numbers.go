package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0
	for !done(l1, l2, carry) {
		digit, nextCarry := addDigits(val(l1), val(l2), carry)
		tail = appendDigit(tail, digit)
		l1, l2, carry = advance(l1), advance(l2), nextCarry
	}
	return dummy.Next
}

func done(l1, l2 *ListNode, carry int) bool {
	return l1 == nil && l2 == nil && carry == 0
}

func val(n *ListNode) int {
	if n == nil {
		return 0
	}
	return n.Val
}

func advance(n *ListNode) *ListNode {
	if n == nil {
		return nil
	}
	return n.Next
}

func addDigits(a, b, carry int) (int, int) {
	sum := a + b + carry
	return sum % 10, sum / 10
}

func appendDigit(tail *ListNode, digit int) *ListNode {
	tail.Next = &ListNode{Val: digit}
	return tail.Next
}
