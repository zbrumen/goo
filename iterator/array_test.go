package iterator

import (
	"context"
	"testing"
)

func TestNewIterator(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	iter := NewIterator[int](data...)
	for iter.HasNext() {
		iter.Next(context.Background())
	}
	if iter.Next(context.Background()) != 0 {
		t.Fail()
	}
}
