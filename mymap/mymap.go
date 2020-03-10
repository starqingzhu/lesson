package mymap

import (
	"fmt"
	"unsafe"
)

/*
字典的key的类型不能为哪些值？
函数、map、切片（原因是：键类型的值之间必须可以施加操作符==和!=，但是这三个不支持。）

什么类型优先考虑做key呢？
求哈希和判等操作的速度越快，对应的类型就越适合作为键类型

什么类型求哈希快呢？
宽度越窄求哈希越快（宽度指占用空间长度）。eg：char、uint8、int8、bool

*/

func init() {
	fmt.Println("---------welcome mymap init----------")
}

func MyMapInit() {
	fmt.Println("---------map init----------")
	mp1 := map[string]int{"1": 1, "2": 2, "3": 3}
	fmt.Printf("mp1->:%+v	len:%d\n", mp1, len(mp1))

	mp2 := make(map[string]int, 10)
	fmt.Printf("mp2->:%+v	len:%d\n", mp2, len(mp2))

	mp3 := map[string]int{}
	fmt.Printf("mp3->:%+v	len:%d\n", mp3, len(mp3))
}

func MyMapSet() {
	fmt.Println("---------map set----------")
	mp1 := make(map[string]int)
	fmt.Printf("初始状态:\nmap1->:%+v	len:%d\n", mp1, len(mp1))
	mp1["1"] = 1
	fmt.Printf("赋值mp1[1] = 1后:\nmap1->:%+v	len:%d\n", mp1, len(mp1))
}

func MyMapGet() {
	fmt.Println("---------map get----------")
	mp1 := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	fmt.Printf("初始状态:\nmap1->:%+v	len:%d\n", mp1, len(mp1))

	value, exist := mp1["1"]
	if exist {
		fmt.Printf("get mp1[1]->%d\n", value)
	} else {
		fmt.Printf("get mp1[1]->fail %v\n", exist)
	}
}

func MyMapRange() {
	fmt.Println("---------map range----------")
	mp1 := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	fmt.Printf("初始状态:\nmap1->:%+v	len:%d\n", mp1, len(mp1))

	fmt.Println("range map(注意顺序，无序):")
	for key, value := range mp1 {
		fmt.Printf("->key:%s	value:%d\n", key, value)
	}
}

func MyMapDel() {
	fmt.Println("---------map range----------")
	mp1 := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	fmt.Printf("初始状态:\nmap1->:%+v	len:%d\n", mp1, len(mp1))

	delete(mp1, "1")

	fmt.Printf("delete 1后:\nmap1->:%+v	len:%d\n", mp1, len(mp1))
}

func TestMyMapFunc() {
	fmt.Println("---------map test func----------")
	mp1 := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	fmt.Printf("初始状态:\nmap1->:%+v	len:%d\n", mp1, len(mp1))

	MyMapFunc(mp1)

	fmt.Printf("delete后:\nmap1->:%+v		len:%d\n", mp1, len(mp1))
}

func MyMapFunc(mp1 map[string]int) {
	fmt.Printf("param pass:\nmap1->:%+v	len:%d\n", mp1, len(mp1))
	delete(mp1, "1")
}

func TestMapCap() {
	fmt.Println("---------map test cap----------")
	mp1 := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	ln, cp := getMapMyInfo(mp1)
	fmt.Printf("初始状态:\nmap1->:%+v	len:%d	cap:%d\n", mp1, ln, cp)
}

type hmap struct {
	count      int
	flags      uint8
	B          uint8
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
}

func getMapMyInfo(m map[string]int) (int, int) {
	point := (**hmap)(unsafe.Pointer(&m))
	value := *point
	return value.count, int(value.B)
	//return int(value.B)

}
