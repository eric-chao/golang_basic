package lee_code

// 有重复计算
func maxArea(height []int) int {
	var length = len(height)
	if length < 2 {
		return 0
	}

	var maxArea int
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			maxArea = max((min(height[i], height[j]))*(j-i), maxArea)
		}
	}

	return maxArea
}
