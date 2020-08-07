package lee_code

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	var base = 1_000_000_000

	const length = 10
	var actualLen = 0
	var arr [length]int
	for i, _ := range arr {
		var b = base / int(math.Pow10(i))
		var d = x / b
		arr[i] = d
		if actualLen == 0 && d != 0 {
			actualLen = length - i
		}
		x = x - d*b
	}

	var index = length - actualLen
	fmt.Printf("arr: %v, index: %d, actualLen: %d \n", arr, index, actualLen)
	for i := 0; i < actualLen/2; i++ {
		var start = index + i
		var end = length - 1 - i
		fmt.Printf("start: %d, val: %d, end: %d, val: %d \n", start, arr[start], end, arr[end])
		if arr[start] != arr[end] {
			return false
		}
	}
	return true
}

func isPalindrome_Better(x int) bool {
	if x < 0 {
		return false
	}
	magic := []int8{}
	for x != 0 {
		magic = append(magic, int8(x%10))
		x /= 10
	}
	fmt.Println(magic)
	for i := 0; i < len(magic)/2; i++ {
		if magic[i] != magic[len(magic)-1-i] {
			return false
		}
	}
	return true
}