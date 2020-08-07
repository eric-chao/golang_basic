package lee_code

import "testing"

func TestReverseInteger(t *testing.T) {
	t.Log(reverse(-12))
	t.Log(reverse(123))
	t.Log(reverse(-12345))
	t.Log(reverse(123456))

	t.Log(reverse(1534236469))
}
