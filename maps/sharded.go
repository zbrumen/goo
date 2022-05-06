package maps

type shardedMap[K comparable, V any] struct {
	shards []Simple[K, V]
	hash   func(key K) int
}

func (s *shardedMap[K, V]) shard(key K) Simple[K, V] {
	return s.shards[s.hash(key)%len(s.shards)]
}

func (s *shardedMap[K, V]) Get(key K) (V, bool) {
	return s.shard(key).Get(key)
}

func (s *shardedMap[K, V]) Set(key K, value V) {
	s.shard(key).Set(key, value)
}

func (s *shardedMap[K, V]) Delete(key K) (V, bool) {
	return s.shard(key).Delete(key)
}

func (s *shardedMap[K, V]) Every(fn func(key K, value V) bool) {
	for _, shard := range s.shards {
		var ok bool
		shard.Every(func(key K, value V) bool {
			if fn(key, value) {
				ok = true
			}
			return ok
		})
		if ok {
			return
		}
	}
}

func (o *shardedMap[K, V]) SetIf(key K, fn func(value V, exists bool) (setValue V, set bool)) (isSet bool) {
	shard := o.shard(key)
	if adv, ok := shard.(Advanced[K, V]); ok {
		return adv.SetIf(key, fn)
	} else {
		val, set := fn(shard.Get(key))
		if set {
			shard.Set(key, val)
		}
		return set
	}
}

func (o *shardedMap[K, V]) DeleteIf(key K, fn func(value V, exists bool) (isDeleted bool)) (isDeleted bool) {
	shard := o.shard(key)
	if adv, ok := shard.(Advanced[K, V]); ok {
		return adv.DeleteIf(key, fn)
	} else {
		isDeleted = fn(shard.Get(key))
		if isDeleted {
			shard.Delete(key)
		}
		return isDeleted
	}
}

func NewShardedMap[K comparable, V any](fn func(key K) int, shards ...Simple[K, V]) Simple[K, V] {
	s := &shardedMap[K, V]{
		hash:   fn,
		shards: shards,
	}
	return s
}

func NewAdvancedShardedMap[K comparable, V any](fn func(key K) int, shards ...Advanced[K, V]) Advanced[K, V] {
	sh := make([]Simple[K, V], len(shards))
	for index, s := range shards {
		sh[index] = s.(Simple[K, V])
	}
	s := &shardedMap[K, V]{
		hash:   fn,
		shards: sh,
	}
	return s
}
