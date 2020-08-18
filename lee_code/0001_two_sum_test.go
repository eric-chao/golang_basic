package lee_code

import "testing"

func TestTwoSum(t *testing.T) {
	var nums, target = []int{2, 11, 15, 7}, 9

	var result = twoSum(nums, target)

	t.Logf("Result 1: %v", result)

	nums, target = []int{3, 2, 4}, 6

	result = twoSum(nums, target)

	t.Logf("Result 2: %v", result)
}

func TestTwoSum2(t *testing.T) {
	var nums, target = []int{2, 11, 15, 7}, 9

	var result = twoSum2(nums, target)

	t.Logf("Result 1: %v", result)

	nums, target = []int{3, 2, 4}, 6

	result = twoSum2(nums, target)

	t.Logf("Result 2: %v", result)
}
