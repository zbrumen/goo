package goo

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func Ok[V any](data V, ok bool) V {
	if ok {
		return data
	}
	panic("must: bool was false")
}

func DebugOk[V any](data V, ok bool) V {
	if ok {
		return data
	}
	_, file, line, _ := runtime.Caller(1)
	cache := strings.Split(file, string(os.PathSeparator))
	panic(fmt.Sprintf("must: bool was false on %s:%d", cache[len(cache)-1], line))
}

func Err[V any](data V, err error) V {
	if err != nil {
		panic(err)
	}
	return data
}

type Resolver[OUT any] struct {
	o OUT
	e error
}

func (r Resolver[OUT]) Then(fn func(o OUT) error) error {
	if r.e != nil {
		return r.e
	}
	return fn(r.o)
}

func Do[OUT any](o OUT, err error) Resolver[OUT] {
	return Resolver[OUT]{
		o, err,
	}
}

func Check[OUT any](o OUT, ok bool) Resolver[OUT] {
	var err error
	if !ok {
		err = errors.New("boolean: false")
	}
	return Resolver[OUT]{
		o, err,
	}
}
