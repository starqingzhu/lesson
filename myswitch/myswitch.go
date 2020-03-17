package myswitch

import (
	"fmt"
)

func init() {
	fmt.Println("---------welcome myswitch init----------")
}

func MySwitchTest() {
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value1[0] {
	case value1[0], value1[1]:
		fmt.Println("myswitch 0 or 1")
		fallthrough
	case value1[2], value1[3]:
		fmt.Println("myswitch 2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("myswitch 4 or 5 or 6")
	}
}
