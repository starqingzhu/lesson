package mystring

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("---------welcome mystring init----------")
}

func TestMystring() {
	var str = "/111/222/333/444"
	var strSep = "/"
	res := strings.Split(str, strSep)
	fmt.Println("mystring res %v", res)

}
