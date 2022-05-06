package iterator

import "context"

type Iterator[V any] interface {
	Next(ctx context.Context) V
	HasNext() bool
}

type Provider[V any] interface {
	Push(ctx context.Context, value V) bool
	Close() error
}

type Chan[V any] interface {
	Iterator[V]
	Provider[V]
}

func Everything[V any](ctx context.Context, i Iterator[V]) []V {
	var out []V
	for i.HasNext() && ctx.Err() == nil {
		out = append(out, i.Next(ctx))
	}
	return out
}
