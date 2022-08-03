package mychan

import (
	"fmt"
	"time"
)

func PrintSelect() {
	c1, c2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- 5
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
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
