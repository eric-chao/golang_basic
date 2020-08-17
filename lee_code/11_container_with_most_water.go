package lee_code

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

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
