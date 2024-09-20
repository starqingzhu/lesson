package mygo

import (
	"fmt"
	"sync"
	"time"
)

/*
协程栈空间：初始值
_StackMin = 2048

协程退出模式（要求优雅）
1.分离模式
创建goroutine者不需要关心其退出，goroutine启动后与其创建者彻底分离，生命周期与其执行的主函数相关，函数返回及goroutine退出。
用途：
	1）一次性任务
	2）常驻后台执行一些特殊任务
2.join模式
用途：
	1）等待一个goroutine退出  返回空消息
	2）获取goroutine的退出状态 chan中返回特定消息结构即可
	3）等待多个goroutine退出
	4）支持超时机制的等待
3.notify-and-wait模式
	1）通知并等待一个goroutine退出
	2）通知并等待多个goroutine退出 sync.waitgroup

退出模式应用

*/

type signal struct {
}

func worker(i int, quit <-chan signal) {
	fmt.Printf("worker %d: is working...\n", i)

LOOP:
	for {
		select {
		default:
			time.Sleep(1 * time.Second)
		case <-quit:
			break LOOP
		}
	}
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(int, <-chan signal), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("worker %d: start to work...\n", i)
			f(i, groupSignal)
			wg.Done()
		}(i + 1)
	}
	go func() {
		wg.Wait()
		c <- signal(struct{}{})
	}()
	return c
}

func PrintGoWork() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	fmt.Println("the group of wokers start to work...")
	time.Sleep(5 * time.Second)

	//通知workers退出
	fmt.Println("notify the group of workers to exis...")
	close(groupSignal)

	<-c
	fmt.Println("the group of workers work done!")

}
