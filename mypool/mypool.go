package mypool

import (
	"fmt"
	"sync"
)

/*
sync.pool
定义：
	临时对象池，它的值可以用来存储临时的对象，结构体类型，它的值被真正使用后，就不应该再被复制了。
	临时对象：不需要持久使用的某一类值。这类值对程序来说可有可无，但如果有的话明显更好。它们的创建和销毁 可以再任何时候发生，并且不影响到程序的功能。
原理
	类型只有两个方法--Put和Get。 Put用于在当前池存放临时对象，接受一个interface{}类型的参数。Get用于从当前池用获取临时对象，返回一个interface{}的值，
	Get方法可能会从当前池中删除掉任何一个值，然后把这个值当做结果返回，如果当前池中没有任何值，这个方法会使用当前池的New字段来创建一个新值，并直接将其返回。
	sync.Pool类型的New字段代表着创建临时对象的函数。即：func() interface{}.该函数的结果值并不会被存入当前的临时对象池中，直接返回给Get方法的调用方。
用途
	帮助程序实现可伸缩性

同步工具：
	context.Context
	sync.WaitGroup
	sync/atomic.Value
	sync.Once
	sync.Pool

*/

func init() {
	fmt.Println("---------welcome mypool init----------")
}

func GetObjPool() sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			return new(int)
		},
	}
}

func TestPool() {
	strPool := &sync.Pool{
		New: func() interface{} { return "" },
	}

	strPool.Put("Hello")
	strPool.Put("World")

	str := strPool.Get().(string)
	fmt.Println(str)

	str += ",Goland!"
	strPool.Put(str)

	str = strPool.Get().(string)
	fmt.Println(str)

	str = strPool.Get().(string)
	fmt.Println(str)

}
