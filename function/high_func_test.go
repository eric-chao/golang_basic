package function

import "testing"

func TestHighFunc(t *testing.T) {
	var f F
	f = func(int int) int {
		return int * 2
	}
	t.Logf("f:%v", higher(12, f))
}
