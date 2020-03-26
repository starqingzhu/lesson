package mymutex

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("---------welcome mymutex init----------")
}

func MyMutexTest() {
	mu := sync.Mutex{}
	mu.Lock()
	mu.Unlock()
}
