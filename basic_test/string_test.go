package basic_test

import "testing"

func TestString(t *testing.T) {
	var str = "abcdefghigklmnopqrstuvwxyz"
	var num = "0123456789"
	var spec = " -+"

	for _, ch := range str {
		t.Log(ch)
	}

	for _, n := range num {
		t.Log(n)
	}

	for _, s := range spec {
		t.Log(s)
	}
}
