package main

import (
    "Treiber-Stack/stack"
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Num of elements = LCM(2,3,4,...,12) * 3000
// 12 is runtime.GOMAXPROCS(0) on my pc
const elements = 83_160_000

func BenchSingleThread() {
    myStack := stack.NewStack[int]()
    startTime := time.Now()
    for i := range elements {
        myStack.Push(i)
    }
    fmt.Println("No goroutines, single thread:", time.Since(startTime).Truncate(time.Millisecond))
}

func BenchMultipleGoroutines() {
    myStack := stack.NewStack[int]()
    wg := sync.WaitGroup{}
    wg.Add(elements)
    startTime := time.Now()
    for i := range elements {
        go func() {
            myStack.Push(i)
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println(elements, "goroutines:", time.Since(startTime).Truncate(time.Millisecond))
}

func BenchNotManyGoroutines() {
    for j := 1; j <= 12; j++ {
        runtime.GOMAXPROCS(j)
        numOfGoroutines := runtime.GOMAXPROCS(0)
        myStack := stack.NewStack[int]()
        wg := sync.WaitGroup{}
        wg.Add(elements)
        startTime := time.Now()
        for range numOfGoroutines {
            go func() {
                for j := range elements / numOfGoroutines {
                    myStack.Push(j)
                    wg.Done()
                }
            }()
        }
        wg.Wait()
        fmt.Println(numOfGoroutines, "goroutines:", time.Since(startTime).Truncate(time.Millisecond))
    }
}
