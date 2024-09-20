package mydefer

import "fmt"

func deferTest() {
	deferTestNoDelay()
	deferTestDelay()
}

func deferTestNoDelay() {
	fmt.Println("==============deferTestNoDelay===============")
	var a = 1
	defer fmt.Println(a)

	a = 2
	return
}

func deferTestDelay() {
	fmt.Println("================deferTestDelay===============")
	var arr = [3]int{1, 2, 3}
	defer printTest(&arr)

	arr[0] = 4
	return
}

func printTest(arr *[3]int) {
	for i := range arr {
		fmt.Println((arr[i]))
	}
}

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func TestDefer() {
	defer func1()
	defer func2()
	defer func3()
}
