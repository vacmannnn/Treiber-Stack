package backoffStack

import (
	"errors"
	"sync/atomic"
)

type state int

const (
	empty state = iota
	busy
	waiting
)

type exchanger[T any] struct {
	item atomic.Value
}

type item[T any] struct {
	value     T
	itemState state
}

func newExchanger[T any]() exchanger[T] {
	result := exchanger[T]{}
	var nilVal T
	result.item.Store(item[T]{nilVal, empty})
	return result
}

func (e exchanger[T]) exchange(value T, duration int) (T, error) {
	var nilVal T
	for i := 0; i < duration; i++ {
		exchangerItem := e.item.Load().(item[T])
		switch exchangerItem.itemState {
		case empty:
			oldItem := item[T]{nilVal, empty}
			newItem := item[T]{value, waiting}
			if e.item.CompareAndSwap(oldItem, newItem) {
				for j := i; j < duration; j++ {
					exchangerItem := e.item.Load().(item[T])
					if exchangerItem.itemState != busy {
						continue
					}
					newItem := item[T]{nilVal, empty}
					e.item.Store(newItem)
					return exchangerItem.value, nil
				}
				return nilVal, errors.New("timeout")

			}
		case waiting:
			oldItem := item[T]{exchangerItem.value, waiting}
			newItem := item[T]{value, busy}
			if e.item.CompareAndSwap(oldItem, newItem) {
				return exchangerItem.value, nil
			}
		case busy:
			break
		}
	}
	return nilVal, errors.New("timeout")
}
