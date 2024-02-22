package stack

import (
    "sync"
    "testing"
)

func TestPop(t *testing.T) {
    myStack := NewStack[int]()
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
    const elements = 123
    myStack := NewStack[int]()
    for i := 0; i < elements; i++ {
        myStack.Push(i)
    }

    if elements != myStack.Size() {
        t.Errorf("Expected %d elements, but got %d", elements, myStack.Size())
    }
}

func TestPushConcurrently(t *testing.T) {
    const goroutineNumber = 123
    myStack := NewStack[int]()
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

// Num of elements = LCM(2,3,4,...,12) * 3000
// 12 is runtime.GOMAXPROCS(0) on my pc
const elements = 83_160_000

func BenchmarkStack_PushNoGoroutines(b *testing.B) {
    for i := 0; i < b.N; i++ {
        myStack := NewStack[int]()
        for i := 0; i < elements; i++ {
            myStack.Push(i)
        }
    }
}

func BenchmarkStack_PushGoroutines(b *testing.B) {
    for i := 0; i < b.N; i++ {
        myStack := NewStack[int]()
        wg := sync.WaitGroup{}
        wg.Add(elements)
        for i := 0; i < elements; i++ {
            go func() {
                myStack.Push(i)
                wg.Done()
            }()
        }
        wg.Wait()
    }
}

func BenchmarkStack_PushSmartGoroutines(b *testing.B) {
    const numOfGoroutines = 6
    for i := 0; i < b.N; i++ {
        myStack := NewStack[int]()
        wg := sync.WaitGroup{}
        wg.Add(elements)
        for i := 0; i < numOfGoroutines; i++ {
            go func() {
                for j := 0; j < elements/numOfGoroutines; j++ {
                    myStack.Push(j)
                    wg.Done()
                }
            }()
        }
        wg.Wait()
    }
}
