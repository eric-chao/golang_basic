package lee_code

import "math"

func reverse(x int) int {

	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}

	var magic = []int8{}
	var index = 0
	for x != 0 {
		var m = int8(x % 10)
		x /= 10
		magic = append(magic, m)
		index++
	}

	var r = 0
	var length = len(magic)
	for _, m := range magic {
		r = r + (int(m) * int(math.Pow10(length -1)))
		length--
	}

	if r < math.MinInt32 || r > math.MaxInt32 {
		return 0
	}

	return r
}


func reverse_better(x int) int {
	if x-int(int32(x)) > 0 { //overflow
		return 0
	}
	flag := 1
	ans := x % 10
	if x < 0 {
		x = -x
		ans = x % 10
		flag = -1
	}

	x = x / 10
	for ; x > 0; x = x / 10 {
		ans = ans*10 + x%10
		if ans != int(int32(ans)) { //overflow
			return 0
		}
	}
	ans = ans * flag
	if ans != int(int32(ans)) { //overflow
		return 0
	}
	return ans
}