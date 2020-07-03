package helper

import "testing"

func TestRandom(t *testing.T) {
	for i := range Random(100) {
		t.Log(i)
	}
}