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

func maxArea2(height []int) int {
	var length = len(height)
	if length < 2 {
		return 0
	}

	var maxArea int
	var left, right = 0, length -1
	for left < right {
		var w = right - left
		var h = 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		var area = w * h
		maxArea = max(maxArea, area)
	}
	return maxArea
}