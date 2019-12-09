package main

import (
	"fmt"
	"lession/hello"
	"lession/myarray"
	"lession/mymap"
	"lession/myslice"
	"lession/mytype"
)

func main() {
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

}
