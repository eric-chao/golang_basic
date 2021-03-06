package lee_code

import "testing"

func TestMaxArea(t *testing.T) {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	var maxArea = maxArea(height)
	t.Logf("MaxArea 1: %d", maxArea)
}

func TestMaxArea2(t *testing.T) {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	var maxArea = maxArea2(height)
	t.Logf("MaxArea 2: %d", maxArea)
}
