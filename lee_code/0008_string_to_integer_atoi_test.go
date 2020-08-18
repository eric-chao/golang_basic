package lee_code

import (
	"math"
	"strconv"
	"testing"
)

func TestMyAtoi(t *testing.T)  {

	var str = "42"

	var val = myAtoi(str)

	t.Log(val)

	var big = "20000000000000000000"

	t.Log(myAtoi(big))

	t.Log(len(strconv.Itoa(math.MaxInt32)))
	t.Log(len(strconv.Itoa(math.MinInt32)))
}