package myatomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
sync/atomic包中提供了几种原子操作？可操作的数据类型又有哪些？
	sync/atomic包中的函数可以做的原子操作有：加法（add）、比较并交换（compare and swap，简称 CAS）、
	加载（load）、存储（store）和交换（swap）。这些函数针对的数据类型并不多。但是，对这些类型中的每一个，
	sync/atomic包都会有一套函数给予支持。这些数据类型有：int32、int64、uint32、uint64、uintptr，以及
	unsafe包中的Pointer。不过，针对unsafe.Pointer类型，该包并未提供进行原子加法操作的函数。
*/
func init() {
	fmt.Println("---------welcome myatomic init----------")
}

var shareVar1 uint32 = 10

func TestMyAtomic() {
	step := 1
	sh2 := atomic.AddUint32(&shareVar1, uint32(step))
	fmt.Printf("myatomic atomic.AddUint32 %d + %d = %d\n", sh2, step, shareVar1)

	var v1 atomic.Value
	v1.Load()
	v1.Store(step)
	fmt.Printf("myatomic atomic.Value %v\n", v1)

	var v2 sync.Once
	v2.Do(func() {})
}
