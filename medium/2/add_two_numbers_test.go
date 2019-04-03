package addTwoNumbers

import "testing"

// var l1_0 = ListNode{
// 	Val:  3,
// 	Next: nil,
// }

// var l1_1 = ListNode{
// 	Val:  4,
// 	Next: &l1_0,
// }

// var l1_2 = ListNode{
// 	Val:  2,
// 	Next: &l1_1,
// }

// var l2_0 = ListNode{
// 	Val:  4,
// 	Next: nil,
// }

// var l2_1 = ListNode{
// 	Val:  6,
// 	Next: &l2_0,
// }

// var l2_2 = ListNode{
// 	Val:  5,
// 	Next: &l2_1,
// }

var l1_1 = ListNode{
	Val:  1,
	Next: nil,
}

var l1_2 = ListNode{
	Val:  8,
	Next: &l1_1,
}

var l2_2 = ListNode{
	Val:  0,
	Next: nil,
}

func TestAddTwoNumbers(t *testing.T) {
	// a := 902
	// t.Log("%%%%%%%: ", a%10, 9/10)
	t.Log(addTwoNumbers(&l1_2, &l2_2))
}
