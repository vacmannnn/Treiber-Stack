package stack

import (
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

func (s Stack) Push(value int) {
    newNode := &node{value: value}
    for {
        head := atomic.LoadPointer(&s.head)
        newNode.next = head
        if atomic.CompareAndSwapPointer(&head, newNode.next, unsafe.Pointer(newNode)) {
            return
        }
    }
}

func (s Stack) Pop() {

}

func (s Stack) Stringer() {

}
