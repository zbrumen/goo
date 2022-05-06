package array

type simpleArray[V any] []*V

func (a *simpleArray[V]) resize(length int) {
	if a.Len() == length {
		return
	}
	var newA = make([]*V, length)
	copy(newA, *a)
	*a = newA
}

func (a *simpleArray[V]) calculateRealSize() int {
	for k := a.Len(); k >= 0; k-- {
		if (*a)[k] != nil {
			return k
		}
	}
	return 0
}

func (a *simpleArray[V]) Len() int {
	return len(*a)
}

func (a *simpleArray[V]) Get(index int) (v V, ok bool) {
	if a.Len() > index {
		entry := (*a)[index]
		ok = entry != nil
		if ok {
			v = *entry
		}
	}
	return
}

func (a *simpleArray[V]) Set(index int, value V) {
	if a.Len() <= index {
		a.resize(index + 1)
	}
	(*a)[index] = &value
}

func (a *simpleArray[V]) Delete(index int) (v V, ok bool) {
	if a.Len() > index {
		(*a)[index] = nil
		ok = true
		if a.Len() == index-1 {
			a.resize(a.calculateRealSize())
		}
	} else {
		ok = false
	}
	return
}

func (a *simpleArray[V]) Append(values ...V) (newLength int) {
	for _, val := range values {
		*a = append(*a, &val)
	}
	return a.Len()
}

func New[V any](values ...V) Simple[V] {
	var pointers simpleArray[V]
	for _, val := range values {
		pointers = append(pointers, &val)
	}
	return &pointers
}
