package gotchas

import "testing"

func TestSlice(t *testing.T) {
	var s1 = make([]int, 5)
	var s2 = make([]int, 5, 8)
	t.Logf("The length of s1: %d\n", len(s1))
	t.Logf("The capacity of s1: %d\n", cap(s1))
	t.Logf("The value of s1: %d\n", s1)

	t.Logf("The length of s1: %d\n", len(s2))
	t.Logf("The capacity of s1: %d\n", cap(s2))
	t.Logf("The value of s1: %d\n", s2)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	t.Logf("The length of s4: %d\n", len(s4))
	t.Logf("The capacity of s4: %d\n", cap(s4))
	t.Logf("The value of s4: %d\n", s4)
}
