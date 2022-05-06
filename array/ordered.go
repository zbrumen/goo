package array

import (
	"github.com/zbrumen/goo/generics"
	"sort"
)

type ordered[V any] struct {
	values Simple[V]
	less   func(i V, iok bool, j V, jok bool) bool
}

func (o *ordered[V]) Delete(index int) (V, bool) {
	return o.values.Delete(index)
}

func (o *ordered[V]) Less(i, j int) bool {
	iV, iOk := o.values.Get(i)
	jV, jOk := o.values.Get(j)
	return o.less(iV, iOk, jV, jOk)
}

func (o *ordered[V]) Swap(i, j int) {
	iV, iOk := o.Get(i)
	jV, jOk := o.Get(j)
	if iOk {
		o.Set(j, iV)
	} else {
		o.Delete(j)
	}
	if jOk {
		o.Set(i, jV)
	} else {
		o.Delete(i)
	}
}

func (o *ordered[V]) Len() int {
	return o.values.Len()
}

func (o *ordered[V]) Get(index int) (v V, ok bool) {
	return o.values.Get(index)
}

func (o *ordered[V]) Set(index int, value V) {
	o.values.Set(index, value)
}

func (o *ordered[V]) Append(values ...V) (newLength int) {
	newLength = o.values.Append(values...)
	sort.Sort(o)
	return newLength
}

func NewCustomOrdered[V any](fn func(i V, iok bool, j V, jok bool) bool, data Simple[V]) Ordered[V] {
	return &ordered[V]{
		less:   fn,
		values: data,
	}
}

func NewOrdered[V generics.Ordered](data ...V) Ordered[V] {
	return NewCustomOrdered(func(i V, iok bool, j V, jok bool) bool {
		return i < j
	}, New[V](data...))
}
