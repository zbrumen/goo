package maps

import (
	"time"
)

type cacheEntry[K comparable, V any] struct {
	name      K
	isDeleted bool
	validFor  time.Time
}

type cacheAdvanced[K comparable, V any] struct {
	internal Advanced[K, V]
	entries  []*cacheEntry[K, V]
	index    Simple[K, *cacheEntry[K, V]]

	validDuration time.Duration
}

func (c *cacheAdvanced[K, V]) Get(key K) (v V, ok bool) {
	return c.internal.Get(key)
}

func (c *cacheAdvanced[K, V]) Set(key K, value V) {
	c.internal.SetIf(key, func(val V, ok bool) (V, bool) {
		entry := &cacheEntry[K, V]{
			name:      key,
			validFor:  time.Now().Add(c.validDuration),
			isDeleted: false,
		}
		c.index.Set(key, entry)
		c.entries = append(c.entries, entry)
		return value, true
	})
}

func (c *cacheAdvanced[K, V]) Delete(key K) (V, bool) {
	var val V
	return val, c.internal.DeleteIf(key, func(value V, ok bool) bool {
		val = value
		if ok {
			entry, _ := c.index.Delete(key)
			if entry != nil {
				entry.isDeleted = true
			}
		}
		return ok
	})
}

func (c *cacheAdvanced[K, V]) Every(f func(key K, value V) bool) {
	c.internal.Every(f)
}

func (c *cacheAdvanced[K, V]) SetIf(key K, fn func(value V, ok bool) (V, bool)) bool {
	return c.internal.SetIf(key, func(value V, ok bool) (V, bool) {
		value, ok = fn(value, ok)
		if ok {
			entry := &cacheEntry[K, V]{
				name:      key,
				validFor:  time.Now().Add(c.validDuration),
				isDeleted: false,
			}
			c.index.Set(key, entry)
			c.entries = append(c.entries, entry)
		}
		return value, ok
	})
}

func (c *cacheAdvanced[K, V]) DeleteIf(key K, fn func(value V, ok bool) bool) bool {
	return c.internal.DeleteIf(key, func(value V, ok bool) bool {
		if fn(value, ok) {
			entry, _ := c.index.Delete(key)
			if entry != nil {
				entry.isDeleted = true
			}
			return true
		}
		return false
	})
}

func CacheAdvanced[K comparable, V any](m Advanced[K, V], timeout time.Duration) Advanced[K, V] {
	return &cacheAdvanced[K, V]{
		internal:      m,
		index:         NewSyncMap[K, *cacheEntry[K, V]](),
		validDuration: timeout,
	}
}
