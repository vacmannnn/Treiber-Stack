package stack

import (
    "fmt"
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

func NewStack() Stack {
    return Stack{}
}

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

// err if pop on empty stack
func (s *Stack) Pop() (int, error) {
    return 0, nil
}

// Top return false if stack is empty
func (s *Stack) Top() (int, bool) {
    if s.head == nil {
        return 0, false
    }
    head := *(*node)(s.head)
    return head.value, true
}

// empty if 0 elements else num of elem
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
