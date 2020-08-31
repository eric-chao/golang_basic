package stucture

import "testing"

func TestSingleLinkList(t *testing.T) {
	singleLink := NewSingleLink()

	arr := [10]Elem{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, e := range arr {
		singleLink.Insert(0, e)
	}

	//
	singleLink.Traverse()

	fiveElem := singleLink.Get(5)
	t.Logf("The 5th elem is %v", fiveElem)
}

func TestLinkSwap(t *testing.T) {
	var head = NewSingleLink()
	var s = []int{1, 2, 3, 4, 5}
	for i, e := range s {
		head.Insert(i, Elem(e))
	}
	var n = head
	for n != nil {
		t.Logf("Val: %+v", n.E)
		n = n.Next
	}
}
