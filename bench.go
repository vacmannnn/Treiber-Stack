package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Num of elements = LCM(2,3,4,...,12) * 3000
// 12 is runtime.GOMAXPROCS(0) on my pc
const elements = 83_160_000

type stackInt[T any] interface {
	Pop() (T, error)
	Push(val T) error
	String() string
}

func BenchSingleThread(myStack stackInt[int]) {
	startTime := time.Now()
	for i := range elements {
		_ = myStack.Push(i)
	}
	fmt.Println("No goroutines, single thread:", time.Since(startTime).Truncate(time.Millisecond))
}

func BenchMultipleGoroutines(myStack stackInt[int]) {
	wg := sync.WaitGroup{}
	wg.Add(elements)
	startTime := time.Now()
	for i := range elements {
		go func() {
			_ = myStack.Push(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(elements, "goroutines:", time.Since(startTime).Truncate(time.Millisecond))
}

func BenchNotManyGoroutines(myStack stackInt[int]) {
	for j := 1; j <= 12; j++ {
		runtime.GOMAXPROCS(j)
		numOfGoroutines := runtime.GOMAXPROCS(0)
		wg := sync.WaitGroup{}
		wg.Add(elements)
		startTime := time.Now()
		for range numOfGoroutines {
			go func() {
				for j := range elements / numOfGoroutines {
					_ = myStack.Push(j)
					wg.Done()
				}
			}()
		}
		wg.Wait()
		fmt.Println(numOfGoroutines, "goroutines:", time.Since(startTime).Truncate(time.Millisecond))
	}
}

func BenchPopAndPush(cStack stackInt[int]) {
	curElems := elements / 2
	for i := 0; i < curElems; i++ {
		_ = cStack.Push(123)
	}
	fmt.Println(&cStack)
	wg := sync.WaitGroup{}
	threads := 4
	wg.Add(threads * 2)
	now := time.Now()
	for i := 0; i < threads; i++ {
		go func() {
			for j := 0; j < curElems/threads; j++ {
				_ = cStack.Push(123)
			}
			wg.Done()
		}()
	}
	fmt.Println(&cStack)
	for i := 0; i < threads; i++ {
		go func() {
			for j := 0; j < curElems/threads; j++ {
				_, _ = cStack.Pop()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("time:", time.Since(now).Truncate(time.Millisecond), "threads:", threads)
	fmt.Println(&cStack)
}
