package hooks

import "context"

type Hooks struct {
	hooks []Hook
}

type Hook interface {
	Run(context.Context)
}

type HookFunc func(context.Context)

func (hf HookFunc) Run(ctx context.Context) {
	hf(ctx)
}

func New() *Hooks {
	return &Hooks{}
}

func (h *Hooks) Add(hook Hook) {
	h.hooks = append(h.hooks, hook)
}

func (h *Hooks) Run(ctx context.Context) {
	for _, h := range h.hooks {
		h.Run(ctx)
	}
}
