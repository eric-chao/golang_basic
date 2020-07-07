package id_gen

import (
	"github.com/bwmarrin/snowflake"
	"testing"
)

func TestGenSnowFlakeId(t *testing.T) {
	n, err := snowflake.NewNode(1)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		t.Logf("id = %d", id.Int64())
		t.Logf("node: %v, step: %v, time: %v", id.Node(), id.Step(),id.Time())
	}
}