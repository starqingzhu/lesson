package myreflect

/*
什么是反射？
	是指计算机程序在运行时可以访问、检测和修改它本身状态或行为的能力

reflect介绍
	反射使用reflect包实现，reflect包实现了运行时反射能力，能够让程序操作不同的对象。
	建立在类型系统之上，它与空接口interface{}密切相关
	每个interface{}类型变量包含一对值（type/value）
	reflect.TypeOf
	reflect.ValueOf

反射三定律
	1.反射是从接口值到反射对象
	2.从反射对象可以获取接口值
	3.要修改反射对象的值，其值必须可以设置


*/
