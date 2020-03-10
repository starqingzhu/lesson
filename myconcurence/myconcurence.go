package myconcurence

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
并发:指让某个函数独立于其他函数执行的能力

调度器:管理且为所有goroutine分配执行执行时间。将操作系统的线程与语言运行时的逻辑处理器绑定，并在逻辑处理器上运行goroutine。调度器在任何
给定的时间，都会全面控制哪个goroutine要在哪个处理器上运行。

并发同步模型:基于通信顺序模型，channel用于goroutine间同步消息


进程+线程模型：操作系统会在物理处理器上调度线程来运行，go语言运行会在逻辑处理器上调度goroutine，每个逻辑处理器都会单独绑定到单个操作系统
线程上。

操作系统线程、逻辑处理器、本地运行队列 关系：
一个goroutine准备并执行，调度器会把其放到一个全局的队列中。之后调度器会把队列中的goroutine分配给逻辑处理器，逻辑处理器将其放到本地队列中，
本地队列中的goroutine会一直等待直到自己被分配到逻辑处理器去执行。

逻辑处理器默认10000个

并发和并行：
并发是指同时管理很多事情，同一时刻并不都执行。并行是指同一时刻不同代码段在多个物理处理器上同时执行。


竞争状态:
如果多个goroutine在没有相互同步的情况下，访问某个共享资源，并试图同时读写这个资源，就处于相互竞争的状态。竟态资源读写需要原子化处理。
检测竟态go自带的工具 race 。eg：go build -race


锁住共享资源的方式:
原子函数 atomic.LoadInt64、atomic.StoreInt64、atomic.AddInt64

互斥锁 sync.Mutex(同一时刻只有一个goroutine进入临界区，直到调用unlock函数其他协程才能进入临界区。当强制推出当前线程后，调度器会再次分配
这个goroutine继续运行)


通道:






*/

const (
	RoutineNum = 2
)

func init() {
	fmt.Println("---------welcome myconcurence init----------")
}

func TestMyGoroutine() {
	fmt.Println("---------test my goroutine----------")

	var wg sync.WaitGroup
	wg.Add(RoutineNum)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println(" ")
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println(" ")
	}()

	fmt.Println("Waiting To Finish")

	wg.Wait()

	fmt.Println("Stop Goroutines")

}

var (
	counter int64
	wgam    sync.WaitGroup
	mutex   sync.Mutex
)

const (
	RoutineAtomicNum = 2
)

func TestMyGoroutineAtomic() {
	fmt.Println("---------test my goroutine atomic----------")
	wgam.Add(RoutineAtomicNum)

	counter = 0
	go incCounter(1)
	go incCounter(2)

	wgam.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wgam.Done()

	fmt.Printf("welcome into incCounter,id:%d\n", id)

	var max int64 = 2
	for counterI := int64(0); counterI < max; counterI++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func TestMyGoroutineMutex() {
	fmt.Println("---------test my goroutine mutex----------")
	wgam.Add(RoutineAtomicNum)

	counter = 0
	go incCounterMutex(1)
	go incCounterMutex(2)

	wgam.Wait()
	fmt.Println("Final Counter:", counter)

}

func incCounterMutex(id int) {
	defer wgam.Done()
	fmt.Printf("welcome into incCounterMutex,id:%d\n", id)

	var max int64 = 2
	for counterI := int64(0); counterI < max; counterI++ {

		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value

		}
		mutex.Unlock()

	}
}

func TestMyGoroutineChan() {
	fmt.Println("---------test my goroutine chan-----------")
	ch1 := make(chan string, 10)
	ch1 <- "hello"

	value := <-ch1

	fmt.Printf("chan value:%s\n", value)
}
