package benchmark

import (
	"fmt"
	tools "lesson/tools/mytime"
	"testing"
)

/*
基准测试辅助函数
*/
func RenameFileTest() {
	path := "./log/"
	file := "hello.log"

	reName := tools.RenameFile(file, tools.GetFileDateName_Day)

	fmt.Sprintf("Rename file,str=%s\n", path+reName)
}

/*
基准测试都在下面
*/

func BenchmarkSprintf(b *testing.B) {
	number := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}
func BenchmarkLogf(b *testing.B) {
	number := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.Logf("%d", number)
	}
}

func BenchmarkRename(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RenameFileTest()
	}
}
