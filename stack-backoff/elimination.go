package stack

import "math/rand"

type eliminationArray[T any] struct {
	duration  int
	exchanger []exchanger[T]
}

func newEliminationArray[T any](cap int) eliminationArray[T] {
	arr := eliminationArray[T]{cap, []exchanger[T]{}}
	arr.exchanger = make([]exchanger[T], arr.duration)
	for i := 0; i < arr.duration; i++ {
		arr.exchanger[i] = newExchanger[T]()
	}
	return arr
}

func (e eliminationArray[T]) visit(value T, rng int) (T, error) {
	slot := rand.Intn(rng)
	return e.exchanger[slot].exchange(value, e.duration)
}
