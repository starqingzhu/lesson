package myrecover

/*
问题一：
从 panic 被引发到程序终止运行的大致过程是什么？
	某个函数中的某行代码有意或无意地引发了一个 panic。这时，初始的 panic 详情会被建立起来，并且该程序的控制权会立即从此行代码转移至调用其所属函数的那行代码上，也就是调用栈中的上一级。
	这也意味着，此行代码所属函数的执行随即终止。紧接着，控制权并不会在此有片刻的停留，它又会立即转移至再上一级的调用代码处。控制权如此一级一级地沿着调用栈的反方向传播至顶端，也就是我们编写的最外层函数那里。
	这里的最外层函数指的是go函数，对于主 goroutine 来说就是main函数。但是控制权也不会停留在那里，而是被 Go 语言运行时系统收回。随后，程序崩溃并终止运行，承载程序这次运行的进程也会随之死亡并消失。与此同时，在这个控制权传播的过程中，panic 详情会被逐渐地积累和完善，并会在程序终止之前被打印出来。
问题二：
对于fmt包下的各种打印函数来说，error类型值的Error方法与其他类型值的String方法是等价的，它们的唯一结果都是string类型的。
 */



import (
	"errors"
	"fmt"
)

func init() {
	recover()
	fmt.Println("---------welcome myrecover init----------")
}

func MyRecoverTest() {
	fmt.Println("myrecover Enter function main.")
	defer func(){
		fmt.Println("myrecover Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("myrecover panic: %s\n", p)
		}
		fmt.Println("myrecover Exit defer function.")
	}()
	// 引发panic。
	panic(errors.New("something wrong"))
	fmt.Println("myrecover Exit function main.")
}