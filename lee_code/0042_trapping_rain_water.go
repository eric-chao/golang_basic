package lee_code

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var sum int
	var left, right, maxLeft, maxRight = 0, len(height) - 1, 0, 0

	for left <= right {
		var r int
		if maxLeft < maxRight {
			r = maxLeft - height[left]
			maxLeft = max(maxLeft, height[left])
			left++
		} else {
			r = maxRight - height[right]
			maxRight = max(maxRight, height[right])
			right--
		}

		if r > 0 {
			sum = sum + r
		}
	}
	return sum
}
