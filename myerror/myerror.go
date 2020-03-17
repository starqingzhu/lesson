package myerror

/*
对于具体错误的判断，Go 语言中都有哪些惯用法？
	1）对于类型在已知范围内的一系列错误值，一般使用类型断言表达式或类型switch语句来判断；
	2）对于已有相应变量且类型相同的一系列错误值，一般直接使用判等操作来判断；
	3）对于没有相应变量且类型未知的一系列错误值，只能使用其错误信息的字符串表示形式来做判断。
*/

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("---------welcome myerror init----------")
}

func MyErrorTest() {
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("myerror request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("myerror error: %s\n", err.Error())
			continue
		}
		fmt.Printf("myerror response: %s\n", resp)
	}
}

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("myerror empty request")
		return
	}
	response = fmt.Sprintf("myerror echo: %s", request)
	return
}
