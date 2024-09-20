package myinterface

//定义新类型
type IntNew int

//类型别名，等价原类型
type Int64New = int64

func Equel() {
	var i64New Int64New
	var i64 int64
	if i64 == i64New {

	}

	var iNew IntNew
	var i int
	if IntNew(i) == iNew {

	}
}
