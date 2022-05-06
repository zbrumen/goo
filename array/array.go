package array

import (
	"sort"
	"sync"
)

type Simple[V any] interface {
	Len() int
	Get(index int) (V, bool)
	Set(index int, value V)
	Delete(index int) (V, bool)
	Append(values ...V) (newLength int)
}

type Synced[V any] interface {
	Simple[V]
	Locker() sync.RWMutex
}

type Ordered[V any] interface {
	Simple[V]
	sort.Interface
}

func Range[V any](a Simple[V], callback func(data V) bool) {
	for k := 0; k < a.Len(); k++ {
		if val, ok := a.Get(k); ok {
			if callback(val) {
				return
			}
		}
	}
}
