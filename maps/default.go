package maps

type defaultMap[K comparable, V any] struct {
	m Simple[K, V]

	defaultValue func() V
}

func (d *defaultMap[K, V]) Get(key K) (V, bool) {
	val, ok := d.m.Get(key)
	if !ok {
		return d.defaultValue(), true
	}
	return val, true
}

func (d *defaultMap[K, V]) Set(key K, value V) {
	d.m.Set(key, value)
}

func (d *defaultMap[K, V]) Delete(key K) (V, bool) {
	return d.m.Delete(key)
}

func (d *defaultMap[K, V]) Every(f func(key K, value V) bool) {
	d.m.Every(f)
}

func WithSimpleDefault[K comparable, V any](m Simple[K, V], defaultValue func() V) Simple[K, V] {
	return &defaultMap[K, V]{
		m:            m,
		defaultValue: defaultValue,
	}
}

type defaultMapAdvanced[K comparable, V any] struct {
	m Advanced[K, V]

	defaultValue func() V
}

func (d *defaultMapAdvanced[K, V]) SetIf(key K, fn func(value V, exists bool) (setValue V, set bool)) (isSet bool) {
	return d.m.SetIf(key, fn)
}

func (d *defaultMapAdvanced[K, V]) DeleteIf(key K, fn func(value V, exists bool) (isDeleted bool)) (isDeleted bool) {
	return d.m.DeleteIf(key, fn)
}

func (d *defaultMapAdvanced[K, V]) Get(key K) (V, bool) {
	val, ok := d.m.Get(key)
	if !ok {
		return d.defaultValue(), true
	}
	return val, true
}

func (d *defaultMapAdvanced[K, V]) Set(key K, value V) {
	d.m.Set(key, value)
}

func (d *defaultMapAdvanced[K, V]) Delete(key K) (V, bool) {
	return d.m.Delete(key)
}

func (d *defaultMapAdvanced[K, V]) Every(f func(key K, value V) bool) {
	d.m.Every(f)
}

func WithAdvancedDefault[K comparable, V any](m Advanced[K, V], defaultValue func() V) Advanced[K, V] {
	return &defaultMapAdvanced[K, V]{
		m:            m,
		defaultValue: defaultValue,
	}
}

func DefaultValue[V any](v V) func() V {
	return func() V {
		return v
	}
}
