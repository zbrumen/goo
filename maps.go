package goo

func Get[K comparable, V any](data map[K]V, key K) (v V, ok bool) {
	v, ok = data[key]
	return
}

func Convert[V any](data any) (v V, ok bool) {
	v, ok = data.(V)
	return v, ok
}
