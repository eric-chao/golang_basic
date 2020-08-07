package pratice

import (
	"testing"
)

func TestGetReward(t *testing.T) {
	getAward(10, []int{})
}

func TestGetAllTotal (t *testing.T) {
	t.Log(getAllTotal(1))
	t.Log(getAllTotal(2))
	t.Log(getAllTotal(3))
	t.Log(getAllTotal(63))
}

func TestMultipleFormula(t *testing.T) {
	getMulFormula(8, []int{})
}
