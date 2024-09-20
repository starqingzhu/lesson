package mymutex

import (
	"fmt"
	"sync"
	"time"
)

/*
条件变量的作用？
	条件变量并不是被用来保护临界区和共享资源的，它是用于协调想要访问共享资源的那些线程的。
	当共享资源的状态发生变化时，它可以被用来通知被互斥锁阻塞的线程

条件变量怎样与互斥锁配合使用？
	条件变量的初始化离不开互斥锁，并且它的方法有的也是基于互斥锁的。
	条件变量提供的方法有三个：等待通知（wait）、单发通知（signal）和广播通知（broadcast）。
*/

var mailbox uint8
var lock sync.RWMutex
var sendCond *sync.Cond = nil
var recvCond *sync.Cond = nil

var WG sync.WaitGroup

func init() {
	fmt.Println("---------welcome mycond init----------")
	sendCond = sync.NewCond(&lock)
	recvCond = sync.NewCond(lock.RLocker())
}

func SendFunc(num int) {
	defer WG.Done()

	for i := 0; i < num; i++ {
		lock.Lock()
		fmt.Println("mycond SendFunc lock...")
		for mailbox == 1 {
			fmt.Println("mycond SendFunc wait...", num)
			sendCond.Wait()
		}
		mailbox = 1
		lock.Unlock()
		fmt.Println("mycond SendFunc unlock...")
		recvCond.Broadcast()
		//recvCond.Signal()
		fmt.Println("mycond SendFunc recvCond Broadcast...")
	}

}

func RecvFunc(goNum int) {
	defer WG.Done()
	WG.Add(1)
	lock.RLock()
	fmt.Println("mycond RecvFunc lock...", goNum)
	for mailbox == 0 {
		fmt.Println("mycond RecvFunc wait...", goNum)
		recvCond.Wait()
	}
	mailbox = 0
	lock.RUnlock()
	fmt.Println("mycond RecvFunc unlock...", goNum)
	sendCond.Signal()
	fmt.Println("mycond RecvFunc sendCond Signal...", goNum)
}

func RecvFuncArr(arrLen int) {
	for i := 0; i < arrLen; i++ {
		go RecvFunc(i)
	}
}

func TestMyCond() {
	const num = 4
	fmt.Println("---------mycond TestMyCond----------")
	WG.Add(1)
	RecvFuncArr(num)
	time.Sleep(500 * time.Millisecond)
	go SendFunc(num)
}
