package myarray

import (
	"fmt"
)

func MyArrInit() {

	/*
		数组几种方式:
		[n]type{}
		[n]type{1,2}
		[n]type{1:11,3:12}
		[...]type{1,2,3}

	*/
	fmt.Println("---------arr init----------")
	var arr0 [5]int64
	fmt.Println("只声明:", arr0)

	arr1 := [5]int64{1, 2, 3}
	fmt.Println("快速创建数组并初始化:", arr1)

	arr2 := [5]int64{1, 2, 3, 4, 5}
	fmt.Println("快速创建数组并初始化:", arr2)

	arr3 := [5]int64{1: 11, 3: 13}
	fmt.Println("快速创建数组并初始化:", arr3)

	arr4 := [...]int64{1, 2}
	fmt.Println("快速创建数组并初始化:", arr4)
}

func MyArrSet() {
	fmt.Println("---------arr 内部 set----------")
	arr0 := [5]int64{}
	fmt.Println(arr0, ":初始状态")

	arr0[0] = 1
	fmt.Println(arr0, ":arr0[0] = 1后")
}

func MyArrsSet() {
	fmt.Println("---------arr 间 set----------")
	var arr1 [5]int64
	var arr2 = [5]int64{1, 2}
	fmt.Println(arr1, ":arr1")
	fmt.Println(arr2, ":arr2")

	arr1 = arr2
	fmt.Println("arr1 = arr2后-------------")
	fmt.Println(arr1, ":arr1")
	fmt.Println(arr2, ":arr2")
}

func MyMutiArr() {
	fmt.Println("---------arr多维数组----------")
	arr := [3][2]int64{{1}, {3}, {5}}
	fmt.Printf("多维数组初始化:\n%+v\n", arr)
	fmt.Printf("访问arr[2][1]:\n%d\n", arr[2][1])

	arr[2][1] = 6
	fmt.Printf("多维数组arr[2][1] = 6后:\n%+v\n", arr)
	fmt.Printf("访问arr[2]:\n%+v\n", arr[2])
}

func MyFuncArr(arr *[5]int64) {
	fmt.Printf("函数间传递 数组:\n%v:*arr\n", *arr)
}

func TestMyFuncArr() {
	var arr = [5]int64{1, 2}
	fmt.Printf("%v:arr\n", arr)
	MyFuncArr(&arr)
}
