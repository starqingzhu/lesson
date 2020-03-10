package mysrclib

/*
iota 在常量声明区有特殊的作用：
1:让编译器为每个常量复制相同的表达式，直到声明区结束（或者遇到一个新的赋值语句）。
2：iota初始值0，之后iota值在每次处理为常量后，都会自增1。
特别声明（每个常量区iota都是从0开始，每过一个变量iota都会1，不管有没有用iota赋值


log包是多协程安全的

*/

import (
	"fmt"
	"log"
)

const (
	MB int = iota
	MC     = 1
	MD     = iota
)

const (
	MA int = iota
	ME
	MX
	MY
	MF
	MG
)

const (
	Prefix = ""
	Flags  = log.Ldate | log.Ltime | log.Lshortfile
)

func init() {
	fmt.Println("---------welcome mylog init----------")
	log.SetPrefix(Prefix)
	log.SetFlags(Flags)
}

func TestLogFmt() {
	fmt.Printf("-----------test logfmt ------------\n")
	fmt.Printf("log iota %d %d %d\n", MB, MC, MD)
}
