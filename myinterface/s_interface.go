package myinterface

import "fmt"

type (
	Interface interface {
		M1()
		M2()
	}
	T struct {
		Interface
	}
	S struct {
	}
)

func (*S) M1() {
	fmt.Printf("S m1...\n")
}
func (*S) M2() {
	fmt.Printf("S m2...\n")
}

func (T) M1() {
	fmt.Printf("T m1...\n")
}

/*
s方法实现 先调用s的方法
s方法未实现，去找接口实现方法
*/
func PrintSInterface() {
	//DumpMethodSet((*T)(nil))
	//var t T
	//var pt *T
	//
	//DumpMethodSet(&t)
	//DumpMethodSet(&pt)
	t := &T{
		Interface: &S{},
	}

	t.M1()
	t.M2()
}
