package timer

import (
	"fmt"
	"time"
)

/*
time.NewTimer
timer.Stop
timer.Reset
time.After

执行一次
*/

func WaitChannel(conn <-chan string) bool {
	timerVal := time.NewTimer(1 * time.Second)
	select {
	case <-conn:
		timerVal.Stop()
	case <-timerVal.C: //超时
		println("waitchannel timeout")
		return false
	}

	return true
}

func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("Delayed 5s, start to do something")
	}
}
