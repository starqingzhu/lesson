package mychan

import (
	"fmt"
	"time"
)

/*
在Go语言中，select和chan是两个重要的并发特性，它们通常一起使用，用来协调不同goroutine之间的通信，实现高效的IO多路复用和异步协作的功能。下面介绍几种常见的使用场景：

多路IO复用：select可以同时监听多个channel的数据情况，从而实现非阻塞IO操作。例如，我们可以利用select同时监听多个网络socket或文件，以响应任意一个IO事件。这种方式适用于需要同时处理多个异步IO事件的场景。

生产者-消费者模型：生产者将数据写入channel中，消费者从channel中读取数据进行处理。使用select能够实现消费者同时监听多个生产者的数据，并选择其中一个可用的channel进行读取。这种方式适用于生产者和消费者之间有多个并行通道的场景。

信号和超时处理：通过将channel和定时器结合使用，可以实现某些延时和超时操作。例如，如果向一个channel中发送一个信号，在另一个goroutine中使用select进行监听，并在超时或收到信号时进行相应的处理。这种方式适用于需要处理超时、取消或中断操作的场景。

并发控制：通过使用带缓冲的channel或无缓冲的channel，并结合select的特性，可以实现对并发任务的控制。例如，可以使用一个带缓冲的channel来限制最大并发数，当channel已满时，阻止新的任务进入，直到某个任务完成后再释放缓冲区。

总之，select和chan是Go语言中非常重要的并发特性，能够极大地提高程序的并发性能和可维护性。虽然这些特性在使用时有一定的复杂性和注意事项，但只要正确使用，并遵循一些设计原则和最佳实践，就可以有效地应对多种场景和需求。
*/

func PrintSelect() {
	c1, c2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- 5
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 7
		close(c2)
	}()

	//var ok1, ok2 bool
	for {
		select {
		case x, ok := <-c1:
			if !ok {
				c1 = nil
			} else {
				fmt.Println(x)
			}

		case x, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				fmt.Println(x)
			}
		}
		if c1 == nil && c2 == nil {
			fmt.Printf("all chan close ok\n")
			break
		}
	}
	fmt.Println("program end...")
}

func worker(c <-chan int) {
	heartBeat := time.NewTicker(time.Second * 30)
	defer heartBeat.Stop()

	for {
		select {
		case <-c:
		case <-heartBeat.C:
			//
		}
	}
}
