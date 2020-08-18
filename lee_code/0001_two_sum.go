package lee_code

// 时间换空间
func twoSum(nums []int, target int) []int {

	var length = len(nums)
	for i, v := range nums {
		for j := i + 1; j < length; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// 空间换时间
func twoSum2(nums []int, target int) []int {

	var valueMap = make(map[int]int)
	for i, v := range nums {
		if _, ok := valueMap[v]; ok {
			return []int{valueMap[v], i}
		}
		valueMap[target-v] = i
	}

	return []int{}
}
