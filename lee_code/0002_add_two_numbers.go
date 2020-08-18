package lee_code

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 没正确理解题意
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var v1 = getNumber(1, l1)
	var v2 = getNumber(1, l2)

	var sum = v1 + v2

	var nums = splitNum(sum)

	var l = &ListNode{}
	for _, n := range nums {
		point := l
		for point.Next != nil {
			point = point.Next
		}

		newNode := &ListNode{Val: n}
		point.Next = newNode
	}
	return l.Next
}

func getNumber(i int, l *ListNode) int {

	if l == nil {
		return 0
	}

	return i*l.Val + getNumber(i*10, l.Next)
}

func splitNum(num int) []int {

	if num == 0 {
		return []int{0}
	}

	var base = 10

	var nums []int
	for num > 0 {
		n := num - num/base*base

		fmt.Printf("n: %d \n", n)
		nums = append(nums, n)
		num = num / 10
	}

	return nums
}
