package gotchas

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSyncPool(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	
	for i := range [1024]int8{} {
		p.Put(i)
	}

	time.Sleep(time.Second)
	for i := 0; i < 1024; i++ {
		t.Log(p.Get())
	}

}

func TestSwitch(t *testing.T) {
	runtime.Gosched()
}
