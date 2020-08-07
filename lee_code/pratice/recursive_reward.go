package pratice

import (
	"fmt"
	"github.com/eric-chao/golang_basic/helper"
	"github.com/thoas/go-funk"
	"math"
)

var rewards = [4]int{1, 2, 5, 10}

func getAward(totalReward int, list []int) {

	if totalReward < 0 {
		return
	}

	if totalReward == 0 {
		fmt.Println(list)
		return
	}

	for _, reward := range rewards {
		var l = make([]int, len(list))
		if len(list) > 0 {
			copy(l, list)
			//var i = copy(l, list)
			//fmt.Printf(" %d copied", i)
		}
		l = append(l, reward)
		getAward(totalReward-reward, l)
	}
}

/*
	一个整数可以被分解为多个整数的乘积。
	例如，6 可以分解为 2x3。请使用递归编程的方法，为给定的整数 n，找到所有可能的分解（1 在解中最多只能出现 1 次）。
	例如，输入 8，输出是可以是 1x8, 8x1, 2x4, 4x2, 1x2x2x2, 1x2x4, ……
*/
func getMulFormula(num int, list []int) {

	if num < 0 {
		return
	}

	if num == 1 {
		if !funk.Contains(list, 1) {
			list = append(list, 1)
		}
		fmt.Println(list)
		return
	}

	for i := range helper.N(num) {
		var val = i + 1
		if (val == 1 && !funk.Contains(list, val)) ||
			(val != 1 && num%val == 0) {
			var l = make([]int, len(list))
			if len(list) > 0 {
				copy(l, list)
			}
			l = append(l, val)
			getMulFormula(num/val, l)
		}
	}
}

func getAllTotal(num int) int64 {

	if num <= 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	return int64(math.Pow(2, float64(num-1))) + getAllTotal(num-1)
}
