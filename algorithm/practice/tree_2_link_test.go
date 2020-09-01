package practice

import (
	"fmt"
	"testing"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToLink(n *Node, link *Node) {
	if n == nil {
		return
	}

	var tmp = link
	for tmp.Right != nil {
		tmp = tmp.Right
	}

	tmp.Right = &Node{Val: n.Val}
	fmt.Printf("#Root %d\n", n.Val)

	if n.Left != nil {
		fmt.Printf("#Left %d\n", n.Left.Val)
		treeToLink(n.Left, link)
	}

	if n.Right != nil {
		fmt.Printf("#Right %d\n", n.Right.Val)
		treeToLink(n.Right, link)
	}
}

func TestConvert(t *testing.T) {
 	var n = &Node{
 		Val: 1,
 		Left: &Node{
 			Val: 2,
 			Left: &Node{
 				Val: 3,
 				Right: &Node{
 					Val: 5,
 					Left: &Node{
 						Val: 4,
					},
 					Right: nil,
				},
			},
		},
		Right: &Node{
 			Val: 6,
			Left: nil,
			Right: nil,
		},
	}
	t.Log("--------------")
	var link = &Node{}
	treeToLink(n, link)
	for link != nil {
		t.Log(link.Val)
		link = link.Right
	}
}
