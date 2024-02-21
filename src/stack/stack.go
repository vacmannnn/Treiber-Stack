package stack

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    "sync/atomic"
    "unsafe"
)

type node struct {
    value int
    next  unsafe.Pointer
}

type Stack struct {
    head unsafe.Pointer
}

// NewStack creates new stack instance
func NewStack() Stack {
    return Stack{}
}

// Push value to top of stack. Concurrency-safety,
// possible to use with many goroutines.
func (s *Stack) Push(value int) {
    newNode := &node{value: value}
    for {
        head := atomic.LoadPointer(&s.head)
        newNode.next = head
        if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(newNode)) {
            return
        }
    }
}

// Pop removes value from top of stack. Returns removed value or
// error if stack was empty.
//
// Concurrency-safety, possible to use with many goroutines.
func (s *Stack) Pop() (int, error) {
    for {
        head := s.head
        if head == nil {
            return 0, errors.New("Pop on empty stack")
        }
        n := *(*node)(head)
        if atomic.CompareAndSwapPointer(&s.head, head, n.next) {
            return n.value, nil
        }
    }
}

// Top returns last element in stack. Returns false if stack was empty
func (s *Stack) Top() (int, bool) {
    if s.head == nil {
        return 0, false
    }
    head := *(*node)(s.head)
    return head.value, true
}

// String describes how many elements on stack, returns
// "empty stack" or "N elements in stack"
func (s *Stack) String() string {
    if s.head == nil {
        return "Empty stack"
    }
    elemCounter := 0
    curHead := s.head
    for curHead != nil {
        head := *(*node)(curHead)
        elemCounter++
        curHead = head.next
    }
    return fmt.Sprintf("%d elements in stack", elemCounter)
}

// Size returns number of elements in stack
func (s *Stack) Size() int {
    elementsInStack, _ := strconv.Atoi(strings.Fields(s.String())[0])
    return elementsInStack
}
