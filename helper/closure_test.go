package helper

import "testing"

func TestClosure1(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer func() {
			t.Logf("index: %d", i)
		}()
	}
}

func TestClosure2(t *testing.T) {
	for i := 0; i < 5; i++ {
		i := i
		defer func() {
			t.Logf("index: %d", i)
		}()
	}
}
func TestClosure3(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			t.Logf("index: %d", i)
		}(i)
	}

}
