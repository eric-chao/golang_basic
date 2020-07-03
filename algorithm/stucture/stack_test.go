package stucture

import "testing"

func TestStack(t *testing.T) {
	items := &[]string{"a", "b", "c", "d", "e", "f"}
	s := NewStack(5)

	for _, e := range *items {
		s.Push(e)
	}

	length := len(*items)
	for i :=0; i< length; i++ {
		t.Logf("[pop] %d, %v", i, s.Pop())
	}
}
