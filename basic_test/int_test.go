package basic

import (
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	var base = 1_000_000_000
	var d = 2147483647 / base

	for i := 0; i < 10; i++ {
		var did = int(math.Pow10(i))
		t.Logf("i: %d, val: %d, base: %d", i, did, base/did)
	}
	t.Log(d)
}

func TestMathMod(t *testing.T) {
	var n1 = 123
	var n2 = -123

	t.Logf("n1 mod = %d, %d", n1%10, n1/10)
	t.Logf("n2 mod = %d, %d", n2%10, n2/10)
}

func TestInt(t *testing.T) {
	var n = 2147483647
	var base = 1_000_000_000

	const length = 10
	var actualLen = 0
	var arr [length]int
	for i, _ := range arr {
		var b = base / int(math.Pow10(i))
		var d = n / b
		arr[i] = d
		if d != 0 {
			actualLen++
		}
		n = n - d*b
	}

	var index = length - actualLen - 1
	for i := 0; i < actualLen/2; i++ {
		if arr[index+1] != arr[length-1-i] {
			//return false
		}
	}

	t.Log(arr)
	t.Logf("length: %d", length)

}
