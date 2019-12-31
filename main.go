package main

import (
	"fmt"
	"lession/hello"
	"lession/myarray"
	"lession/mymap"
	"lession/myslice"
	"lession/mytype"
	"time"
)

func main() {
	defer durationTotal(time.Now().UnixNano())
	fmt.Println("hello测试专用----------------------------------->>>>>>>>>>>")
	hello.PrintHello()

	fmt.Println("数组测试专用----------------------------------->>>>>>>>>>>")
	myarray.MyArrInit()
	myarray.MyArrSet()
	myarray.MyArrsSet()
	myarray.MyMutiArr()
	myarray.TestMyFuncArr()

	fmt.Println("切片测试专用----------------------------------->>>>>>>>>>>")
	myslice.MySliceInit()
	myslice.MySliceSet()
	myslice.MySliceAppend()
	myslice.MySliceRange()
	myslice.MyMutiSlice()
	myslice.TestMyFuncSlice()
	myslice.MySliceMemInfo()

	fmt.Println("映射测试专用----------------------------------->>>>>>>>>>>")
	mymap.MyMapInit()
	mymap.MyMapSet()
	mymap.MyMapGet()
	mymap.MyMapRange()
	mymap.MyMapDel()
	mymap.TestMyMapFunc()
	mymap.TestMapCap()

	fmt.Println("语言类型测试专用----------------------------------->>>>>>>>>>>")
	mytype.Init()
	mytype.Func()
	mytype.TestInterfaceNotify()
	mytype.TestMutiInterfaceNotify()
	mytype.TestInsertion()
	testPrivateValue()
	testPrivateInsertion()

}


func durationTotal(begin int64){
	duration := time.Now().UnixNano() - begin
	fmt.Printf("duration elapse: %dus\n",duration/1000000)

}

func testPrivateValue(){
	fmt.Println("---------test private value----------")
	counter := mytype.New(10)
	fmt.Printf("counter:%d\n",counter)
}

func testPrivateInsertion(){
	fmt.Println("---------test private insertion----------")
	ad1 := mytype.Admin1{
		Level: 0,
	}
	ad1.Name = "sunbin01"
	//ad1.email = "starqingzhu@163.com" //未公开变量无法外部访问

	ad2 := mytype.Admin2{
		Level: 1,
	}
	ad2.Name = "sunbin02"
	ad2.Email = "starqingzhu@163.com"

	fmt.Printf("ad1:\n\t%+v\nad2:\n\t%+v\n",ad1,ad2)
}
