package basic

import (
	"sync/atomic"
	"testing"
)

var map1 = map[string]int{"key1": 1, "key2": 2, "key3": 3}
var map2 = map[string]int{"key4": 4, "key5": 5, "key6": 6}

func TestAtomicValue(t *testing.T) {

	var v = atomic.Value{}
	v.Store(readOnly{m: map1})
	read, _ := v.Load().(readOnly)

	if val, ok := read.m["key1"]; ok {
		t.Logf("key1 = %d", val)
	}

	v.Store(readOnly{m: map2})

	read, _ = v.Load().(readOnly)
	if val, ok := read.m["key1"]; ok {
		t.Logf("key1 = %d", val)
	}

	if val, ok := read.m["key3"]; ok {
		t.Logf("key3 = %d", val)
	}

	if val, ok := read.m["key5"]; ok {
		t.Logf("key5 = %d", val)
	}

}

type readOnly struct {
	m       map[string]int
	amended bool // true if the dirty map contains some key not in m.
}
