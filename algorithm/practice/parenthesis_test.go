package practice

import (
	"fmt"
	"testing"
)

/*
	40
	41
	91
	93
	123
	125
*/
func TestParenthesis(t *testing.T) {
	var str = "abcd"
	fmt.Println(str[:len(str)-1])
	var s1 = "()[]]{}"

	fmt.Println("--- --- ")

	var signMap = make(map[int]int, 3)
	signMap[40] = 41
	signMap[91] = 93
	signMap[123] = 125

	var s []int
	for _, ch := range s1 {
		var c = int(ch)
		if _, ok := signMap[int(ch)]; ok {
			s = append(s, c)
		} else {
			var l = len(s)
			if l == 0 {
				fmt.Println(false)
				return
			}

			if r, ok := signMap[s[l-1]]; ok {
				if r == c {
					fmt.Println("###")
					s = s[:l-1]
				}
			}

		}
		fmt.Println(s)
		fmt.Println(" --- --- ")
	}

	if len(s) > 0 {
		fmt.Println(false)
		return
	}

	fmt.Println(true)
}
