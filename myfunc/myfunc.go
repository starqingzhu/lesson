package myfunc

/*
函数作用：
封装代码、分割功能、解藕逻辑

函数签名：


高阶函数：满足以下任意条件都是高阶函数
1 接收其他的函数作为参数传入
2 把其他的函数作为结果返回


闭包：


*/

import (
	"errors"
	"fmt"
)

//函数签名
type (
	Printer func(contents string) (n int, err error)
	Operate func(x, y int) int
	OperateCheck func(x, y int) (int,error)
)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func calculate(x int, y int, op Operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}


func genCalculator(op Operate) (OperateCheck) {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func Testfunc() {
	var p Printer
	p = printToStd
	p("func something")

	op := func(x, y int) int {
		return x * y
	}
	res, err := calculate(1, 2, op)
	if err != nil {
		fmt.Printf("func calculate err %s\n", err)
		return
	}
	fmt.Printf("func calculate res %d\n", res)

	checkRes := genCalculator(op)
	res,err = checkRes(2,2)
	if err != nil{
		fmt.Printf("func checkRes err %d\n", err)
		return
	}
	fmt.Printf("func checkRes res %d\n", res)


}
