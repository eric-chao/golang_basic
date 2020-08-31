package stucture

import "fmt"

type Node struct {
	Val  string
	Next *Node
}

type LinkList struct {
	Length int
	Head   *Node
}

func NewLinkList() *LinkList {
	return &LinkList{}
}

func (s *LinkList) AddHead(v string) {
	var e = &Node{Val: v}
	if s.Head == nil {
		s.Head = e
		s.Length = 1
		return
	}

	e.Next = s.Head
	s.Head = e
	s.Length++
}

func (s *LinkList) AddTail(v string) {
	var e = &Node{Val: v}

	if s.Head == nil {
		s.Head = e
		s.Length = 1
		return
	}

	var n = s.Head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = e
	s.Length++
}

func (s *LinkList) Traverse() {
	var n = s.Head
	for n != nil {
		fmt.Printf(" -> %s \n", n.Val)
		n = n.Next
	}
}
