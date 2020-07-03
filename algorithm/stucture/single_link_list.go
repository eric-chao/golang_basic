package stucture

import "fmt"

type Elem int

type LinkNode struct {
	E    Elem
	Next *LinkNode
}

func NewSingleLink() *LinkNode {
	return &LinkNode{0, nil}
}

//在链表的第i个位置前插入一个元素e，复杂度为o(n)
func (head *LinkNode) Insert(i int, e Elem) bool {
	p := head
	j := 0
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	s := &LinkNode{E: e}
	s.Next = p.Next
	p.Next = s
	return true
}

//遍历链表
func (head *LinkNode) Traverse() {
	point := head.Next
	for nil != point {
		fmt.Println(point.E)
		point = point.Next
	}
	fmt.Println("--------done----------")
}

//删除链表中第i个节点，复杂度为o(n)
func (head *LinkNode) Delete(i int) bool {
	p := head
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}

	p.Next = p.Next.Next
	return true
}

// 获取链表中的第i个元素，复杂度为o(n)
func (head *LinkNode) Get(i int) Elem  {
	p := head.Next
	for j:= 1; j< i ;j++  {
		if nil == p {
			//表示返回错误
			return -100001
		}
		p=p.Next
	}

	return p.E
}