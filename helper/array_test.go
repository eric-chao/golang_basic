package helper

import "testing"

func TestArray(t *testing.T) {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		t.Logf("inner func(), arr = %v", arr)
	}(x)

	t.Logf("outer func(), arr = %v", x)
}

func TestSlice(t *testing.T) {
	x := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		t.Logf("inner func(), arr = %v", arr)
	}(x)

	t.Logf("outer func(), arr = %v", x)
}