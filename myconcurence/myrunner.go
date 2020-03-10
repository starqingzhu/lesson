package myconcurence

import "fmt"

/*
并发模式要解决问题
1.控制程序生命周期
2.管理可复用的资源池
3.创建处理任务的协程池
*/

func init() {
	fmt.Println("---------welcome myrunner init----------")
}

func TestRunner() {
	fmt.Println("---------test my goroutine runner-----------")
}
