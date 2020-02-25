package mychan

import (
	"fmt"
)

/*
带缓冲
不带缓冲
 */

var ch1 = make(chan string,10)

func BufChan(){
	var chanNodeList = []string{"hello","world","work"}
	for _,node := range chanNodeList{
		ch1 <- node
	}

	close(ch1)
}

func TestBufChan(){
	go BufChan()
	//time.Sleep(1*time.Second)

	var res string
	for res = range ch1 {
		fmt.Printf("chan res:%s\n",res)
		continue
	}

}

