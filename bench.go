package main

import (
	"fmt"
	"github.com/vacmannnn/Treiber-Stack/backoffStack"
	"github.com/vacmannnn/Treiber-Stack/stack"
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
		_ = myStack.Push(i)
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
			_ = myStack.Push(i)
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
					_ = myStack.Push(j)
					wg.Done()
				}
			}()
		}
		wg.Wait()
		fmt.Println(numOfGoroutines, "goroutines:", time.Since(startTime).Truncate(time.Millisecond))
	}
}

var threadsNum = [...]int{2, 4, 6, 8, 10, 12}

func BenchPopAndPush() {
	cStack := backoffStack.NewStack[int]()
	abc := 24_000_000
	// bStack := backoffStack.NewStack[int]()
	for i := 0; i < abc/24; i++ {
		cStack.Push(123)
	}
	fmt.Println(&cStack)
	wg := sync.WaitGroup{}
	wg.Add(8)
	now := time.Now()
	for i := 0; i < 4; i++ {
		go func() {
			for j := 0; j < abc/4; j++ {
				cStack.Push(123)
			}
			wg.Done()
		}()
	}
	fmt.Println(&cStack)
	for i := 0; i < 4; i++ {
		go func() {
			for j := 0; j < abc/4; j++ {
				cStack.Pop()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("time:", time.Since(now).Truncate(time.Millisecond))
	fmt.Println(&cStack)
}
