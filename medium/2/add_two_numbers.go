package addTwoNumbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var p *ListNode
	var head *ListNode
	var total int
	carry := 0
	for l1 != nil || l2 != nil || carry == 1 {
		if l1 != nil && l2 != nil {
			total = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			total = l1.Val + carry
			l1 = l1.Next
		} else if l2 != nil {
			total = l2.Val + carry
			l2 = l2.Next
		} else if carry == 1 {
			total = carry
		}
		carry = total / 10

		var q ListNode
		q.Val = total % 10
		q.Next = nil
		if head == nil {
			head = &q
		} else {
			p.Next = &q
		}
		p = &q

	}
	return head
}
