package maps

import (
	"sync"
)

type syncMap[K comparable, V any] struct {
	data map[K]V
	sync sync.RWMutex
}

func (s syncMap[K, V]) Every(fn func(key K, value V) bool) {
	s.sync.RLock()
	defer s.sync.RUnlock()
	for k, v := range s.data {
		if fn(k, v) {
			return
		}
	}
}

func (s *syncMap[K, V]) Get(key K) (V, bool) {
	s.sync.RLock()
	out, ok := s.data[key]
	s.sync.RUnlock()
	return out, ok
}

func (s *syncMap[K, V]) Set(key K, value V) {
	s.sync.Lock()
	s.data[key] = value
	s.sync.Unlock()
}

func (s *syncMap[K, V]) Delete(key K) (V, bool) {
	value, ok := s.Get(key)
	if ok {
		s.sync.Lock()
		delete(s.data, key)
		s.sync.Unlock()
	}
	return value, ok
}

func (s *syncMap[K, V]) SetIf(key K, fn func(value V, ok bool) (V, bool)) bool {
	s.sync.Lock()
	defer s.sync.Unlock()
	value, ok := s.data[key]
	if input, ok := fn(value, ok); ok {
		s.data[key] = input
		return true
	}
	return false
}

func (s *syncMap[K, V]) DeleteIf(key K, fn func(value V, ok bool) bool) bool {
	s.sync.Lock()
	defer s.sync.Unlock()
	value, ok := s.Get(key)
	if fn(value, ok) {
		delete(s.data, key)
		return true
	}
	return false
}

func NewSyncMap[K comparable, V any]() Advanced[K, V] {
	return &syncMap[K, V]{
		data: map[K]V{},
		sync: sync.RWMutex{},
	}
}
