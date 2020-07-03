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
