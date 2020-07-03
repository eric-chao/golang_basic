package helper

import "testing"

const (
	WaitToBeEnabled = iota + 2
	Enabled
	History
)

func TestIotaConst(t *testing.T) {
	t.Logf("WaitToBeEnabled %d", WaitToBeEnabled)
	t.Logf("Enabled %d", Enabled)
	t.Logf("History %d", History)
}