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

func (s *Stack) Pop() {

}

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
