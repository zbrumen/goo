package iterator

import (
	"context"
	"fmt"
	"sync"
)

type chanIterator[V any] struct {
	chs    chan V
	once   sync.Once
	closed bool
}

func (c chanIterator[V]) Push(ctx context.Context, value V) bool {
	select {
	case c.chs <- value:
		return true
	case <-ctx.Done():
		return false
	}
}

func (c chanIterator[V]) Close() error {
	if !c.closed {
		c.once.Do(func() {
			close(c.chs)
			c.closed = true
		})
		return nil
	}
	return fmt.Errorf("closed")
}

func (c chanIterator[V]) Next(ctx context.Context) V {
	var v V
	select {
	case v = <-c.chs:
	case <-ctx.Done():
	}
	return v
}

func (c chanIterator[V]) HasNext() bool {
	return !c.closed
}

func NewChanIterator[V any](buffer int) Chan[V] {
	return chanIterator[V]{
		chs:    make(chan V, buffer),
		once:   sync.Once{},
		closed: false,
	}
}

func Link[V any](buffer int) (Iterator[V], Provider[V]) {
	c := NewChanIterator[V](buffer)
	return c, c
}
