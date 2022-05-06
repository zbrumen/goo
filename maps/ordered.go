package maps

import (
	"context"
	"github.com/zbrumen/goo/array"
	"github.com/zbrumen/goo/generics"
	"github.com/zbrumen/goo/iterator"
	"sort"
)

type orderedMap[K comparable, V any] struct {
	ordered array.Ordered[K]
	m       Simple[K, V]
}

func (o *orderedMap[K, V]) Get(key K) (V, bool) {
	return o.m.Get(key)
}

func (o *orderedMap[K, V]) Set(key K, value V) {
	o.m.Set(key, value)
	o.ordered.Append(key)
	sort.Sort(o.ordered)
}

func (o *orderedMap[K, V]) Delete(key K) (V, bool) {

	return o.m.Delete(key)
}

func (o *orderedMap[K, V]) Every(f func(key K, value V) bool) {
	array.Range[K](o.ordered, func(key K) bool {
		val, ok := o.m.Get(key)
		if ok {
			return f(key, val)
		}
		return false
	})
}

func WithSimpleOrdered[K comparable, V any](ctx context.Context, less func(i K, iOk bool, j K, jOk bool) bool, s Simple[K, V]) (Simple[K, V], error) {
	return &orderedMap[K, V]{
		ordered: array.NewCustomOrdered[K](less, array.New[K](iterator.Everything[K](ctx, Keys[K](ctx, 8, s))...)),
		m:       s,
	}, ctx.Err()
}

func WithGenericOrdered[K generics.Ordered, V any](ctx context.Context, s Simple[K, V]) (Simple[K, V], error) {
	return &orderedMap[K, V]{
		ordered: array.NewOrdered[K](iterator.Everything[K](ctx, Keys[K](ctx, 8, s))...),
		m:       s,
	}, ctx.Err()
}

type orderedAdvancedMap[K comparable, V any] struct {
	ordered array.Ordered[K]
	m       Advanced[K, V]
}

func (o *orderedAdvancedMap[K, V]) SetIf(key K, fn func(value V, exists bool) (setValue V, set bool)) (isSet bool) {
	return o.m.SetIf(key, func(value V, exists bool) (setValue V, set bool) {
		setValue, set = fn(value, exists)
		if set {
			o.ordered.Append(setValue)
		}
		return setValue, set
	})
}

func (o *orderedAdvancedMap[K, V]) DeleteIf(key K, fn func(value V, exists bool) (isDeleted bool)) (isDeleted bool) {
	return o.m.DeleteIf(key, fn)
}

func (o *orderedAdvancedMap[K, V]) Get(key K) (V, bool) {
	return o.m.Get(key)
}

func (o *orderedAdvancedMap[K, V]) Set(key K, value V) {
	o.m.Set(key, value)
	o.ordered.Append(key)
}

func (o *orderedAdvancedMap[K, V]) Delete(key K) (V, bool) {
	return o.m.Delete(key)
}

func (o *orderedAdvancedMap[K, V]) Every(f func(key K, value V) bool) {
	array.Range[K](o.ordered, func(key K) bool {
		val, ok := o.m.Get(key)
		if ok {
			return f(key, val)
		}
		return false
	})
}

func WithAdvancedOrdered[K comparable, V any](ctx context.Context, less func(i K, iOk bool, j K, jOk bool) bool, s Advanced[K, V]) (Advanced[K, V], error) {
	return &orderedAdvancedMap[K, V]{
		ordered: array.NewCustomOrdered[K](less, array.New[K](iterator.Everything[K](ctx, Keys[K](ctx, 8, s.(Simple[K, V])))...)),
		m:       s,
	}, ctx.Err()
}

func WithAdvancedGenericOrdered[K generics.Ordered, V any](ctx context.Context, s Advanced[K, V]) (Advanced[K, V], error) {
	return &orderedAdvancedMap[K, V]{
		ordered: array.NewOrdered[K](iterator.Everything[K](ctx, Keys[K](ctx, 8, s.(Simple[K, V])))...),
		m:       s,
	}, ctx.Err()
}
