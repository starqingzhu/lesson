package mypool

import (
	"fmt"
	"sync"
	"time"
)

func init() {
	fmt.Println("---------welcome mypool init----------")
}


func TestMyPool(){
	var bytePool = sync.Pool{
		New:func() interface{} {
			b := make([]byte, 1024)
			return &b
		},
	}

	b := time.Now().UnixNano()/1000
	for i := 0; i < 10000; i++{
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().UnixNano()/1000


	fmt.Println("mypool with pool ", c - b, "us")
}
