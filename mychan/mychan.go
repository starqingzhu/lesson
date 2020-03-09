package mychan

import (
	"fmt"
	"math/rand"
)

/*
带缓冲
不带缓冲

通道的基本特性：
1 对于同一个通道，发送操作之间是互斥的，接受操作之间也是互斥的。
2 发送操作中和接收操作中对元素值的处理都是不可分割的（原子操作）
3 发送操作在完全完成之前会被阻塞。接收操作也是如此

通道分类：
单向通道（发送通道 make(chan<- int,1)、接收通道 make(<-chan int,1)）
双向通道 (make(chan, 1)

单向通道意义？
重要用途是约束其他代码行为。（配合接口使用）


*/

var ch1 = make(chan string, 10)

func BufChan() {
	var chanNodeList = []string{"hello", "world", "work"}
	for _, node := range chanNodeList {
		ch1 <- node
	}

	close(ch1)
}

func TestBufChan() {
	go BufChan()
	//time.Sleep(1*time.Second)

	chanVar := <-ch1
	fmt.Printf("chan var:%s\n", chanVar)
	var res string
	for res = range ch1 {
		fmt.Printf("chan res:%s\n", res)
		continue
	}


	// 准备好几个通道。
	chanLen := 1
	intChannels := [3]chan int{
		make(chan int, chanLen),
		make(chan int, chanLen),
		make(chan int, chanLen),
	}
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	select {
	case <-intChannels[0]:
		fmt.Println("chan The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("chan The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("chan The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("chan No candidate case is selected!")
	}
}
