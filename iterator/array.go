package iterator

import (
	"context"
	"github.com/zbrumen/goo/array"
)

type arrayIterator[V any] []V

func (a *arrayIterator[V]) Next(ctx context.Context) V {
	var out V
	if a.HasNext() {
		out = ([]V)(*a)[0]
		*a = ([]V)(*a)[1:]
	}
	return out
}

func (a *arrayIterator[V]) HasNext() bool {
	return len(*a) > 0
}

func NewIterator[V any](data ...V) Iterator[V] {
	return (*arrayIterator[V])(&data)
}

func NewArrayIterator[V any](arr array.Simple[V]) Iterator[V] {
	out := NewChanIterator[V](arr.Len())
	go func() {
		array.Range[V](arr, func(val V) bool {
			out.Push(context.Background(), val)
			return false
		})
		_ = out.Close()
	}()
	return out
}
