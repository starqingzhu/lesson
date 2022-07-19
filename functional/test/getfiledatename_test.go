package test

/*
单元测试：基础测试、表组测试
单元测试：用来测试包或者（程序的一部分代码或者一组代码的）函数
测试目的：确认目标代码在给定的场景下，有没有按照期望工作
测试手段：正向路径测试、负向路径测试
基础测试：只是用一组参数和结果进行测试
表组测试：使用多组参数和结果进行测试


测试用例
文件名以 _test.go结尾
测试函数名要以 Test 开头 eg:TestXxx

测试时要在测试目录下
go test -v -test.run 测试函数名
*/

import (
	"os"
	"testing"
)

func TestRenameFile(t *testing.T) {
	//测试用例并行
	t.Parallel()

	path := "./log/"
	file := "hello.log"

	reName := tools.RenameFile(file, tools.GetFileDateName_Day)

	t.Logf("Rename file,str=%s\n", path+reName)
}
func Test1Env(t *testing.T) {
	//测试用例并行
	t.Parallel()

	key := "GOPATH"
	value := os.Getenv(key)
	t.Logf("key:%s  value:%s\n", key, value)
}
