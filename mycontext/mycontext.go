package mycontext

import (
	"context"
	"fmt"
	"sync/atomic"
)

/*
context使用場景:
	上下文信息传递 （request-scoped），比如处理 http 请求、在请求处理链路上传递信息
	控制子 goroutine 的运行
	超时控制的方法调用
	可以取消的方法调用
*/

func init() {
	fmt.Println("---------welcome mycontext init----------")
}

func addNum(num *int32, i int, f func()) {
	atomic.AddInt32(num, 1)
	f()
}

func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	cxt, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFunc()
			}
		})
	}
	<-cxt.Done()
	fmt.Println("End.")
}

func TestcoordinateWithContext() {
	//coordinateWithContext()
	//ctx := context.Background()
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0002")
	ctx = context.WithValue(ctx, "key3", "0003")
	ctx = context.WithValue(ctx, "key4", "0004")

	fmt.Println(ctx.Value("key1"))
}
