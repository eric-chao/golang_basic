package helper

import (
	"testing"
	"time"
)

func TestMonth(t *testing.T) {
	var now = time.Now()
	t.Logf("Year: %d", now.Year())
	t.Logf("Month: %d", now.Month())
} 
