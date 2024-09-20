package myfunc

import "fmt"

func init() {
	fmt.Printf("-----welcome param change ------\n")
}

/*
变长参数要求：
1.只需要有一个变长参数
2.变长参数必须在最后一个参数

需要注意：
1.传入值需要类型匹配
*/

func ComParams(p ...int64) string {
	var ret string
	for _, c := range p {
		ret += fmt.Sprintf("%d ", c)
	}

	return ret
}

func PrintComParam() {
	var arrInt64 = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ret := ComParams(arrInt64...)
	fmt.Printf("param change ret:%s\n", ret)
}

//模拟函数重载
func conCat(sep string, args ...interface{}) string {
	var ret string
	for _, v := range args {
		switch v.(type) {
		case int, int8, int16, int64, int32:
			ret += fmt.Sprintf("%d", v)
		case string:
			ret += fmt.Sprintf("%s", v)
		default:
			continue
		}
		ret += sep
	}
	return ret
}

func PrintConCat() {
	var sep = " "
	var args = []interface{}{1, "hello", "?"}

	ret := conCat(sep, args...)
	fmt.Printf("concat ret:%s\n", ret)
}
