package stack

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
	"unsafe"
)

type node[T any] struct {
	value T
	next  unsafe.Pointer
}

type Stack[T any] struct {
	head unsafe.Pointer
}

// NewStack creates new stack instance.
func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

// Push value to top of stack.
//
// Concurrency-safety, possible to use with many goroutines.
func (s *Stack[T]) Push(value T) error {
	if s == nil {
		return errors.New("nil pointer to stack")
	}
	newNode := &node[T]{value: value}
	for {
		head := atomic.LoadPointer(&s.head)
		newNode.next = head
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(newNode)) {
			return nil
		}
	}
}

// Pop removes value from top of stack. Returns removed value or
// error if stack was empty.
//
// Concurrency-safety, possible to use with many goroutines.
func (s *Stack[T]) Pop() (T, error) {
	if s == nil {
		var nilVal T
		return nilVal, errors.New("nil pointer to stack")
	}
	for {
		head := s.head
		if head == nil {
			var nilVal T
			return nilVal, errors.New("pop on empty stack")
		}
		n := *(*node[T])(head)
		if atomic.CompareAndSwapPointer(&s.head, head, n.next) {
			return n.value, nil
		}
	}
}

// Top returns last element in stack. Returns false if stack was empty.
func (s *Stack[T]) Top() (T, bool) {
	if s == nil || s.head == nil {
		var nilVal T
		return nilVal, false
	}
	head := *(*node[T])(s.head)
	return head.value, true
}

// String describes how many elements on stack, returns
// "empty stack" or "N elements in stack".
//
// Aware: not concurrency-safety.
func (s *Stack[T]) String() string {
	if s == nil || s.head == nil {
		return "Empty stack"
	}
	elemCounter := 0
	curHead := s.head
	for curHead != nil {
		head := *(*node[T])(curHead)
		elemCounter++
		curHead = head.next
	}
	return fmt.Sprintf("%d elements in stack", elemCounter)
}

// Size returns number of elements in stack.
//
// Aware: not concurrency-safety.
func (s *Stack[T]) Size() int {
	elementsInStack, _ := strconv.Atoi(strings.Fields(s.String())[0])
	return elementsInStack
}
