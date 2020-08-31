package basic

import (
	"context"
	"testing"
)

func f(ctx context.Context) {
	context.WithValue(ctx, "foo", -6)
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	f(ctx)
	t.Log(ctx.Value("foo"))

	c := context.WithValue(ctx, "foo1", 6)
	t.Log(c.Value("foo1"))
}
