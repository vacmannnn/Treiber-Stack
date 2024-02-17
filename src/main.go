package main

import (
    "Treiber-Stack/stack"
    "fmt"
    "runtime"
    "sync"
    "time"
)

const goroutineNumber = 50000000

func main() {
    fmt.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(0))
    myStack := stack.NewStack()
    fmt.Println("Single thread")
    now := time.Now()
    for i := range goroutineNumber {
        myStack.Push(i)
    }
    fmt.Println(time.Now().Sub(now).Truncate(time.Millisecond))
    fmt.Println(&myStack)

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
    fmt.Println(&myStack)
    fmt.Println("\n\n\n")
}
