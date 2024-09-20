package gpool

import (
	"fmt"
	"go.uber.org/atomic"
	"sync"
	"testing"
)

var wg sync.WaitGroup
var counter atomic.Int32

func HandleLogin(fpId int64) {
	fmt.Printf("fpid:%d\n", fpId)
	counter.Add(1)
	wg.Done()
}
func JobWrapper(fpId int64) Job {
	return func() {
		HandleLogin(fpId)
	}
}

func TestGPool(t *testing.T) {
	var numWorks = 10
	var jobQueueLen = 1000
	var jobNum = 10000
	wg.Add(jobNum)
	p := NewPool(numWorks, jobQueueLen)
	p.Start()

	for i := 0; i < jobNum; i++ {
		p.Push(JobWrapper(int64(i)))
	}

	wg.Wait()
	p.Release()

	fmt.Printf("counter:%d\n", counter.Load())
}
