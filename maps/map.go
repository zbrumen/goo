package maps

import (
	"context"
	"github.com/zbrumen/goo/iterator"
)

type Simple[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
	Delete(key K) (value V, edited bool)

	Every(func(key K, value V) (continueFlag bool))
}

type Advanced[K comparable, V any] interface {
	Simple[K, V]

	SetIf(key K, fn func(value V, exists bool) (setValue V, set bool)) (isSet bool)
	DeleteIf(key K, fn func(value V, exists bool) (isDeleted bool)) (isDeleted bool)
}

func Keys[K comparable, V any](ctx context.Context, length int, m Simple[K, V]) iterator.Iterator[K] {
	ch := iterator.NewChanIterator[K](length)
	go func() {
		m.Every(func(key K, value V) bool {
			return !ch.Push(ctx, key)
		})
		_ = ch.Close()
	}()
	return ch
}

func Values[K comparable, V any](ctx context.Context, length int, m Simple[K, V]) iterator.Iterator[V] {
	ch := iterator.NewChanIterator[V](length)
	go func() {
		m.Every(func(key K, value V) bool {
			return !ch.Push(ctx, value)
		})
		_ = ch.Close()
	}()
	return ch
}

type Set[K comparable, V any] struct {
	key   K
	value V
}

func (s Set[K, V]) Key() K {
	return s.key
}

func (s Set[K, V]) Value() V {
	return s.value
}

func Range[K comparable, V any](ctx context.Context, length int, m Simple[K, V]) iterator.Iterator[Set[K, V]] {
	ch := iterator.NewChanIterator[Set[K, V]](length)
	go func() {
		m.Every(func(key K, value V) bool {
			return !ch.Push(ctx, Set[K, V]{
				key:   key,
				value: value,
			})
		})
		_ = ch.Close()
	}()
	return ch
}
