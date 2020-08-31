package stucture

import (
	"fmt"
	"testing"
)

var values = []string{"A", "B", "C", "D", "E", "F", "G"}

func TestAddHead(t *testing.T) {
	var l = NewLinkList()
	for _, e := range values {
		l.AddHead(e)
	}

	l.Traverse()
	t.Log(" --- --- ")
}

func TestAddTail(t *testing.T) {
	var l = NewLinkList()
	for _, e := range values {
		l.AddTail(e)
	}

	l.Traverse()
	t.Log(" --- --- ")
}

func TestLinkSwap2(t *testing.T) {
	var l = NewLinkList()
	for _, e := range values {
		l.AddTail(e)
	}
	l.Traverse()
	fmt.Println("--- ---")

	var r = NewLinkList()
	var n = l.Head
	for n != nil {
		var f = &Node{Val: n.Val}
		var s *Node
		if n.Next == nil {
			s = nil
		} else {
			s = &Node{Val: n.Next.Val}
		}

		if r.Head == nil {
			if s != nil {
				r.Head = s
				r.Head.Next = f
			} else {
				r.Head = f
			}
		} else {
			var h = r.Head
			for h.Next != nil {
				h = h.Next
			}
			if s != nil {
				h.Next = s
				h.Next.Next = f
			} else {
				h.Next = f
			}
		}

		if n.Next == nil {
			break
		}
		n = n.Next.Next
	}
	fmt.Println("--- ---")
	r.Traverse()
}
