package hooks_test

import (
	"context"
	"testing"

	"github.com/lestrrat-go/hooks"
	"github.com/stretchr/testify/assert"
)

func TestHooks(t *testing.T) {
	h := hooks.New()

	expected := make(map[int]struct{})
	for i := 0; i < 10; i++ {
		expected[i] = struct{}{}
		i := i
		h.Add(hooks.HookFunc(func(ctx context.Context) {
			delete(expected, i)
		}))
	}

	h.Run(context.Background())

	if !assert.Len(t, expected, 0) {
		return
	}
}
