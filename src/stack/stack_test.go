package stack

import (
    "fmt"
    "strings"
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
        elementsInStack := strings.Fields(myStack.String())[0]
        t.Errorf("Expected %d elements, but get %s", elements, elementsInStack)
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
    if myStack.String() != fmt.Sprintf("%d elements in stack", elements) {
        elementsInStack := strings.Fields(myStack.String())[0]
        t.Errorf("Expected %d elements, but get %s", elements, elementsInStack)
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
    if myStack.String() != fmt.Sprintf("%d elements in stack", goroutineNumber) {
        elementsInStack := strings.Fields(myStack.String())[0]
        t.Errorf("Expected %d elements, but get %s", goroutineNumber, elementsInStack)
    }
}
