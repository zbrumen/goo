package array

import "sync"

type syncArray[V any] struct {
	data Simple[V]
	sync sync.RWMutex
}

func (s *syncArray[V]) Len() int {
	s.sync.RLock()
	defer s.sync.RUnlock()
	return s.data.Len()
}

func (s *syncArray[V]) Get(index int) (V, bool) {
	s.sync.RLock()
	defer s.sync.RUnlock()
	return s.data.Get(index)
}

func (s *syncArray[V]) Set(index int, value V) {
	s.sync.Lock()
	defer s.sync.Unlock()
	s.data.Set(index, value)
}

func (s *syncArray[V]) Delete(index int) (V, bool) {
	s.sync.Lock()
	defer s.sync.Unlock()
	return s.data.Delete(index)
}

func (s *syncArray[V]) Append(values ...V) (newLength int) {
	s.sync.Lock()
	defer s.sync.Unlock()
	return s.data.Append(values...)
}

func (s *syncArray[V]) Locker() sync.RWMutex {
	return s.sync
}

func WithSyncLocking[V any](data Simple[V]) Synced[V] {
	return &syncArray[V]{
		data: data,
		sync: sync.RWMutex{},
	}
}
