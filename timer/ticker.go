package timer

import (
	"fmt"
	"time"
)

/*
周期性执行
*/

func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Ticker tick.")
		}
	}

	//for range ticker.C {
	//	fmt.Println("Ticker tick.")
	//}
}
