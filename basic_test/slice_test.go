package basic

import "testing"

func TestSliceLenAndCap(t *testing.T) {
	var l = make([]int, 0, 10)
	t.Logf("l len = %d, cap = %d", len(l), cap(l))

	var m = make([]int, 10)
	t.Logf("m len = %d, cap = %d", len(m), cap(m))

}
