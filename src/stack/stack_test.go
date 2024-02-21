package stack

import (
    "sync"
    "testing"
)

func TestPop(t *testing.T) {
    myStack := NewStack()
    elements := 3
    for i := 0; i < elements; i++ {
        myStack.Push(i)
    }
    if myStack.String() != "3 elements in stack" {
        t.Errorf("Expected %d elements, but get %d", elements, myStack.Size())
    }

    for i := 0; i < elements; i++ {
        res, err := myStack.Pop()
        if err != nil {
            t.Error("Unexpected error:", err)
        } else if res != elements-1-i {
            t.Errorf("Expected %d on top of stack, but get %d", elements-1-i, res)
        }
    }

    _, err := myStack.Pop()
    if err == nil {
        t.Error("Stack expected to be empty")
    }
}

func TestPush(t *testing.T) {
    const elements = 100000
    myStack := NewStack()
    for i := 0; i < elements; i++ {
        myStack.Push(i)
    }

    if elements != myStack.Size() {
        t.Errorf("Expected %d elements, but got %d", elements, myStack.Size())
    }
}

func TestPushConcurrently(t *testing.T) {
    const goroutineNumber = 100000
    myStack := NewStack()
    wg := sync.WaitGroup{}
    wg.Add(goroutineNumber)
    for i := 0; i < goroutineNumber; i++ {
        go func(i int) {
            myStack.Push(i)
            wg.Done()
        }(i)
    }
    wg.Wait()

    if goroutineNumber != myStack.Size() {
        t.Errorf("Expected %d elements, but got %d", goroutineNumber, myStack.Size())
    }
}
