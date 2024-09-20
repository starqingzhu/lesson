package mydefer

/*

主函数：
	创建defer的函数
延迟函数：
	defer后面的函数

规则：
	1.延迟函数的参数再defer语句出现时就已经确定
	2.延迟函数执行按后进先出顺序执行，即先出现的defer最后执行
	3.延迟函数可能操作主函数的具名返回值

*/
