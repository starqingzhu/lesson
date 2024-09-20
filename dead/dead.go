package dead

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func DeadLoop() {
	for {
		add(3, 5)
		//fmt.Println("deadloop  ...")
		//time.Sleep(time.Millisecond * 1)
	}
}

func DeadLoop2() {
	for {
		fmt.Println("deadloop 2 ...")
		//time.Sleep(time.Second * 1)
	}
}
