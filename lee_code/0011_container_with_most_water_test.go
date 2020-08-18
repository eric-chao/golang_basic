package lee_code

import "testing"

func TestMaxArea(t *testing.T) {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	var maxArea = maxArea(height)
	t.Logf("MaxArea: %d", maxArea)
}
