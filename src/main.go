package main

import (
    "Treiber-Stack/stack"
    "fmt"
    "runtime"
    "sync"
    "time"
)

const goroutineNumber = 1000000

func main() {
    simplePopTest()
    simplePushBench()
}

func simplePopTest() {
    myStack := stack.NewStack()
    myStack.Push(10)
    myStack.Push(123)
    myStack.Push(900)
    fmt.Println(&myStack)
    res, _ := myStack.Pop()
    fmt.Println(res, &myStack)
    res, _ = myStack.Pop()
    fmt.Println(res, &myStack)
    res, _ = myStack.Pop()
    fmt.Println(res, &myStack)
    res, err := myStack.Pop()
    fmt.Println(res, err, &myStack)
}

func simplePushBench() {
    fmt.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(0))
    myStack := stack.NewStack()
    fmt.Println("Single thread")
    now := time.Now()
    for i := range goroutineNumber {
        myStack.Push(i)
    }
    fmt.Println(time.Now().Sub(now).Truncate(time.Millisecond))
    top, _ := myStack.Top()
    fmt.Println(&myStack, top)

    fmt.Println()
    myStack = stack.NewStack()
    wg := sync.WaitGroup{}
    wg.Add(goroutineNumber)
    fmt.Println("Concurrent/parallel")
    now = time.Now()
    for i := range goroutineNumber {
        go func(i int) {
            myStack.Push(i)
            wg.Done()
        }(i)
    }
    wg.Wait()
    fmt.Println(time.Now().Sub(now).Truncate(time.Millisecond))
    top, _ = myStack.Top()
    fmt.Println(&myStack, top)
    fmt.Println("\n\n\n")
}
