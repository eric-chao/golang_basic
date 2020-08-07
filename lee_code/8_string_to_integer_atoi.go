package lee_code

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {

	if len(str) <= 0 {
		return 0
	}

	var b strings.Builder
	var sign string
	var index = 0
	for _, ch := range str {

		fmt.Println(ch)

		if index == 0 && ch == 32 {
			continue
		}

		if ch < 48 || ch > 57 {
			if index == 0 && (ch == 43 || ch == 45) {
			} else {
				break
			}
		}

		if index == 0 {
			if ch == 43 {
				// sign = "+"
			}
			if ch == 45 {
				sign = "-"
			}
		}

		if ch >= 48 && ch <= 57 {
			b.WriteString(string(ch))
		}

		index++
	}

	if b.Len() > 0 {
		var result = sign + b.String()
		var val, err = strconv.Atoi(result)
		if err != nil {
			fmt.Println(err)
			if e, ok := err.(*strconv.NumError); ok && e.Err == strconv.ErrRange {
				if sign == "" || sign == "+" {
					return math.MaxInt32
				}
				if sign == "-" {
					return math.MinInt32
				}
			}
			return 0
		}

		if val >= math.MaxInt32 {
			return math.MaxInt32
		}

		if val <= math.MinInt32 {
			return math.MinInt32
		}

		return val
	}

	return 0
}

func myAtoi_Better(str string) int {
	var ans int64 = 0
	var sign int64 = 1
	start := false

	for _, b := range str {
		if b <= '9' && b >= '0' {
			if !start {start = true}
			ans = ans * 10 + int64(b) - int64('0')
			if ans * sign > math.MaxInt32 {
				return math.MaxInt32
			} else if ans * sign < math.MinInt32 {
				return math.MinInt32
			}
		} else if !start && b == '+' {
			start = true
		} else if !start && b == '-' {
			start = true
			sign = -1
		} else if !start && b == ' '{
			continue
		} else {
			break
		}
	}
	return int(ans * sign)
}