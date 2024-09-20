package main

import (
	"fmt"
	"lesson/mydefer"
	"time"
)

const (
	MAX_PROC = 2
)

func init() {
	fmt.Println("----------------------------------------------main begin--------------------------------------------------------")
}

func main() {
	// defer durationTotal(time.Now().UnixNano())
	// ret := runtime.GOMAXPROCS(MAX_PROC)
	// fmt.Println("ret", ret)
	mydefer.TestDefer()
	//mygnet.Init()

	// go server.ServerInit(":8081")
	//server.ServerHttpInit()
	//time.Sleep(10000 * time.Second)

	//myreflect.PrintReflect()
	//myreflect.PrintReflect2()

	//mychan.PrintMyLockCount()
	//mychan.PrintSelect()

	//mycontext.TestcoordinateWithContext()
	// mypool.TestPool()
	//go dead.DeadLoop2()
	//go dead.DeadLoop()
	//for {
	//	time.Sleep(time.Second * 1)
	//	fmt.Println("I got scheduled!")
	//}

	//myinterface.PrintInterfaceType()
	//myinterface.PrintSInterface()
	//myinterface.PrintNilInterface()
	//myinterface.PrintEmptyInterface()
	//myinterface.PrintNotEmptyInterface()
	//myinterface.PrintInterfaceInterDetail()
	//myinterface.PrintInterfaceInterCombination()
	//myinterface.PrintInterfaceHandler()
	//
	//myfunc.Testfunc()
	//myfunc.PrintComParam()
	//myfunc.PrintConCat()

	//fmt.Println("mynet测试专用----------------------------------->>>>>>>>>>>")
	//mynet.GetPublicNetAddrTest()
	//
	//fmt.Println("hello测试专用----------------------------------->>>>>>>>>>>")
	//hello.PrintHello()
	//
	//fmt.Println("数组测试专用----------------------------------->>>>>>>>>>>")
	//myarray.MyArrInit()
	//myarray.MyArrSet()
	//myarray.MyArrsSet()
	//myarray.MyMutiArr()
	//myarray.TestMyFuncArr()
	//
	//fmt.Println("切片测试专用----------------------------------->>>>>>>>>>>")
	//myslice.MySliceInit()
	//myslice.MySliceSet()
	//myslice.MySliceAppend()
	//myslice.MySliceRange()
	//myslice.MyMutiSlice()
	//myslice.TestMyFuncSlice()
	//myslice.MySliceMemMyInfo()
	//myslice.MySliceRangeSum()
	//
	// fmt.Println("映射测试专用----------------------------------->>>>>>>>>>>")
	// mymap.Simple1()
	//mymap.MyMapInit()
	//mymap.MyMapSet()
	//mymap.MyMapGet()
	//mymap.MyMapRange()
	//mymap.MyMapDel()
	//mymap.TestMyMapFunc()
	//mymap.TestMapCap()
	//
	//fmt.Println("语言类型测试专用----------------------------------->>>>>>>>>>>")
	//mytype.Init()
	//mytype.Func()
	//mytype.TestInterfaceNotify()
	//mytype.TestMutiInterfaceNotify()
	//mytype.TestInsertion()
	//mytype.TestInsertionFloat()
	//testPrivateValue()
	//testPrivateInsertion()
	//mytype.TestInsert2()
	//
	//fmt.Println("协程测试专用-------------------------------------->>>>>>>>>>>")
	//myconcurence.TestMyGoroutine()
	//myconcurence.TestMyGoroutineAtomic()
	//myconcurence.TestMyGoroutineMutex()
	//myconcurence.TestMyGoroutineChan()
	//myconcurence.TestRunner()
	//
	//fmt.Println("标准库使用测试专用-------------------------------------->>>>>>>>>>>")
	//mysrclib.TestLogFmt()
	//mysrclib.TestMylogger()
	//mysrclib.TestJsonDecodePerson()
	//mysrclib.TestJsonEncodePerson()
	//mysrclib.TestMyIo()
	//
	//fmt.Println("http库使用测试专用-------------------------------------->>>>>>>>>>>")
	//myhttp.TestMyHttp()
	//
	//fmt.Println("chan使用测试专用-------------------------------------->>>>>>>>>>>")
	//mychan.TestBufChan()
	//
	//fmt.Println("sort使用测试专用-------------------------------------->>>>>>>>>>>")
	//mysort.MySort()
	//
	//fmt.Println("func使用测试专用-------------------------------------->>>>>>>>>>>")
	//myfunc.Testfunc()
	//
	//fmt.Println("switch使用测试专用-------------------------------------->>>>>>>>>>>")
	//myswitch.MySwitchTest()
	//
	//fmt.Println("error使用测试专用-------------------------------------->>>>>>>>>>>")
	//myerror.MyErrorTest()
	//
	//fmt.Println("recover使用测试专用-------------------------------------->>>>>>>>>>>")
	//myrecover.MyRecoverTest()
	//
	////fmt.Println("mylog使用测试专用-------------------------------------->>>>>>>>>>>")
	////mylog.TestMylog()
	//
	//fmt.Println("myrandom 使用测试专用-------------------------------------->>>>>>>>>>>")
	//myrandom.TestRandom()
	//
	//fmt.Println("OrderIdTest 使用测试专用-------------------------------------->>>>>>>>>>>")
	//myorder.GetOrderIdTest()
	//
	//fmt.Println("myhttp 使用测试专用-------------------------------------->>>>>>>>>>>")
	//myhttp.TestMyAddrHttp()

	// mypprof.CreateProfileFile()

}

func durationTotal(begin int64) {
	duration := time.Now().UnixNano() - begin
	fmt.Printf("\n\nduration elapse: %dus\n\n\n", duration/1000000)

}

func testPrivateValue() {
	fmt.Println("---------test private value----------")
	//counter := mytype.New(10)
	//fmt.Printf("counter:%d\n", counter)
}

func testPrivateInsertion() {
	fmt.Println("---------test private insertion----------")
	//ad1 := mytype.Admin1{
	//	Level: 0,
	//}
	//ad1.Name = "sunbin01"
	////ad1.email = "starqingzhu@163.com" //未公开变量无法外部访问
	//
	//ad2 := mytype.Admin2{
	//	Level: 1,
	//}
	//ad2.Name = "sunbin02"
	//ad2.Email = "starqingzhu@163.com"
	//
	//fmt.Printf("ad1:\n\t%+v\nad2:\n\t%+v\n", ad1, ad2)
}
