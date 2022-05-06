package generics

import (
	"context"
	"fmt"
	"sync"
)

type Duplex[IN any, OUT any] struct {
	in  <-chan IN
	out chan<- OUT

	once sync.Once
}

func (d Duplex[IN, OUT]) Recv() <-chan IN {
	return d.in
}

func (d Duplex[IN, OUT]) Send(ctx context.Context, data OUT) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	select {
	case d.out <- data:
		return
	case <-ctx.Done():
		return fmt.Errorf("context done")
	}
}

func (d Duplex[IN, OUT]) Close() error {
	d.once.Do(func() {
		close(d.out)
	})
	return nil
}

func NewDuplex[IN any, OUT any](buffer int) (Duplex[IN, OUT], Duplex[OUT, IN]) {
	in := make(chan IN, buffer)
	out := make(chan OUT, buffer)
	return Duplex[IN, OUT]{
			in, out, sync.Once{},
		}, Duplex[OUT, IN]{
			out, in, sync.Once{},
		}
}
