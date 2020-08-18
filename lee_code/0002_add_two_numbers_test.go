package lee_code

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var l1 = &ListNode {
		Val: 2,
		Next: &ListNode {
			Val: 4,
			Next: &ListNode {
				Val: 3,
				Next: nil,
			},
		},
	}

	var l2 = &ListNode {
		Val: 5,
		Next: &ListNode {
			Val: 6,
			Next: &ListNode {
				Val: 4,
				Next: nil,
			},
		},
	}
	
	var l = addTwoNumbers(l1, l2)
	for l != nil {
		t.Logf("Result: %+v", l)
		l = l.Next
	}
}
