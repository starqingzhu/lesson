package myslice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func init() {
	fmt.Println("---------welcome myslice init----------")
}

func MySliceInit() {
	fmt.Println("---------myslice init----------")
	sl1 := make([]int64, 5)
	fmt.Printf("myslice sl1:%v\tlen:%d\tcapacity:%d\n", sl1, len(sl1), cap(sl1))

	sl2 := make([]int64, 3, 5)
	fmt.Printf("myslice sl2:%v\tlen:%d\tcapacity:%d\n", sl2, len(sl2), cap(sl2))

	sl3 := []int64{1, 2, 3, 4, 5}
	fmt.Printf("myslice sl3:%v\tlen:%d\tcapacity:%d\n", sl3, len(sl3), cap(sl3))

	//2种空切片的方式
	//nil切片
	var sl4 []int64
	fmt.Printf("myslice sl4:%v    \tlen:%d\tcapacity:%d 空切片\n", sl4, len(sl4), cap(sl4))
	//fmt.Printf("sl4==nil:%v\n",sl4 == nil)

	//空切片
	sl5 := make([]int64, 0)
	fmt.Printf("myslice sl5:%v    \tlen:%d\tcapacity:%d 空切片\n", sl5, len(sl5), cap(sl5))

	sl6 := []int64{1, 2, 3, 4}
	fmt.Printf("myslice sl6:%v\tlen:%d\tcapacity:%d 空切片\n", sl6, len(sl6), cap(sl6))

	sl7 := sl6[0:3:4]
	fmt.Printf("myslice sl7:%v\tlen:%d\tcapacity:%d 空切片\n", sl7, len(sl7), cap(sl7))

}

func MySliceSet() {
	fmt.Println("---------myslice set----------")
	sl0 := make([]int64, 2, 5)
	fmt.Printf("myslice sl0:%v    \tlen:%d\tcapacity:%d\n", sl0, len(sl0), cap(sl0))
	sl0[0] = 1
	fmt.Printf("myslice sl0[0] = 1后:\nsl0:%v    \tlen:%d\tcapacity:%d\n", sl0, len(sl0), cap(sl0))
}

func MySliceAppend() {
	fmt.Println("---------myslice append----------")
	sl0 := make([]int64, 2, 5)
	fmt.Printf("myslice sl0:%v    \tlen:%d\tcapacity:%d\n", sl0, len(sl0), cap(sl0))
	sl0 = append(sl0, 10)
	fmt.Printf("myslice append(sl0,10)后:\nsl0:%v    \tlen:%d\tcapacity:%d\n", sl0, len(sl0), cap(sl0))

	//fmt(println((*reflect.SliceHeader)(unsafe.Pointer(&sl0))).Data)
	fmt.Printf("\n")
	sl1 := []int64{11, 12}
	sl2 := append(sl0, sl1...)
	fmt.Printf("myslice sl1:%v	\tlen:%d\tcapacity:%d\n", sl1, len(sl1), cap(sl1))
	fmt.Printf("myslice append(sl0,sl1...)后:\nsl2:%v    \tlen:%d\tcapacity:%d\n", sl2, len(sl2), cap(sl2))

}

func MySliceRange() {
	fmt.Println("---------myslice range----------")
	sl := []int64{10, 11, 12, 13, 14}
	for index, value := range sl {
		fmt.Printf("myslice index:%d	value:%d\n", index, value)
	}
	fmt.Println("myslice range 后的value是副本验证:")
	for index, value := range sl {
		fmt.Printf("myslice index:%d	value_addr:%X	ElemAddr:%X\n", index, &value, &sl[index])
	}
	fmt.Println("myslice range 后下标用_忽略:")
	for _, value := range sl {
		fmt.Printf("myslice index:_	value:%d\n", value)
	}
	fmt.Println("myslice 传统for循环:")
	for index := 0; index < len(sl); index++ {
		fmt.Printf("myslice index:%d	sl[index]:%d\n", index, sl[index])
	}
}

func MyMutiSlice() {
	fmt.Println("---------muti myslice----------")
	sl := [][]int{{1, 2}, {3, 4}}
	fmt.Printf("sl:%+v\n", sl)
	for index := 0; index < len(sl); index++ {
		fmt.Printf("myslice index:%d addr:%X\n", index, &sl[index][0])
		//fmt.Printf("index:%d addr:%X\n",index,((*reflect.SliceHeader)(unsafe.Pointer(&sl[index]))).Data)
		fmt.Printf("myslice index:%d sl_addr:%X sl[index]_addr:%X\n", index, ((*reflect.SliceHeader)(unsafe.Pointer(&sl))).Data, ((*reflect.SliceHeader)(unsafe.Pointer(&sl[index]))).Data)
	}

}

func TestMyFuncSlice() {
	sl := []int{1, 2, 3, 4}
	MyFuncSlice(&sl)
}

func MyFuncSlice(sl *[]int) {
	fmt.Println("---------func myslice MyFuncSlice----------")
	for index, value := range *sl {
		fmt.Printf("myslice index:%d	value:%d\n", index, value)
	}
}

func MySliceMemMyInfo() {
	fmt.Println("---------func myslice MySliceMemMyInfo----------")
	//sl := make([][]int, 2,10)
	sl := [][]int{{1, 2, 6, 7}, {3, 4, 5}}
	fmt.Printf("myslice 切片的初始状态:\n%+v\n", sl)

	sl_head_addr := uintptr(unsafe.Pointer(&sl))
	fmt.Printf("myslice sl_head_addr:%X\n", sl_head_addr)

	//ss :=  uintptr(*(*int)(unsafe.Pointer(uintptr(*(*int)(unsafe.Pointer(sl_head_addr))))))

	//64位情况下，指针占8字节
	length := *(*int)(unsafe.Pointer(sl_head_addr + 8))
	fmt.Printf("myslice 最外层切片长度:%d\n", length)

	sl_0_head1_addr := (*int)((unsafe.Pointer(&sl)))
	fmt.Printf("myslice 外层切片head中的ptr指向地址:%X\n", *sl_0_head1_addr)
	fmt.Printf("myslice 切片地址列表:\n%X\n%X\n", uintptr(unsafe.Pointer(&(sl[0]))), uintptr(unsafe.Pointer(&(sl[1]))))

	fmt.Printf("myslice 内层切片ptr值:%X\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&(sl[0]))))))
	fmt.Printf("myslice 内层切片第一个data第一个元素的地址:%X\n", &sl[0][0])

}

func MySliceRangeSum() {
	fmt.Println("---------func myslice MySliceRangeSum ----------")
	sl1 := []int64{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice MySliceRangeSum before: %v\n", sl1)
	for index, value := range sl1 {
		if index == len(sl1)-1 {
			sl1[0] += value
		} else {
			sl1[index+1] += sl1[index]
		}
	}

	fmt.Printf("myslice MySliceRangeSum after: %v\n", sl1)
}

func MySliceChild(){
	fmt.Println("---------func myslice MySliceChild ----------")
	sl1 := []int64{1,2,3,4}
	fmt.Printf("myslice len:%d cap:%d sl1:%v\n",len(sl1), cap(sl1),sl1)
	sl2:= sl1[1:len(sl1)]
	fmt.Printf("myslice len:%d cap:%d sl2:%v\n",len(sl2), cap(sl2),sl2)

}
