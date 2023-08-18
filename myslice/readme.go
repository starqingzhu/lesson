package myslice

/*
slice 底层结构
	unsafe.Pointer
	len
	cap

扩容规则：
	1.如果原slice容量小于1024，则新slice容量扩容为原来2倍
	2.如果原slice容量大于等于1024，则新slice容量扩大为原来的1.25倍


slice copy
	将源切片的数据追个拷贝到目的切片指向的数组中，拷贝数量取两个切片长度的最小值。		注意：copy过程中不会扩容。


*/
