package basic

import (
	"testing"
	"time"
)

func TestSliceLenAndCap(t *testing.T) {
	var l = make([]int, 0, 10)
	t.Logf("l len = %d, cap = %d", len(l), cap(l))

	var m = make([]int, 10)
	t.Logf("m len = %d, cap = %d", len(m), cap(m))

}

// bad demo
func TestSliceLog(t *testing.T) {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			t.Logf("char: %v", v)
		}()
	}
	// select {}    // 阻塞模式
	time.Sleep(time.Second)
}