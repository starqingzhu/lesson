package mystring

/*
string 是什么：
	string是8bit字节的集合，通常但并不一定是UTF-8编码的文本

关注点：
	1.string可以为空（长度为0），但是不会是nil。
	2.string对象不可以修改。

结构：
	type stringStruct struct {
		str unsafe.Pointer		//	字符串的首地址
		len int					//	字符串的长度
	}

[]byte转string
	需要一次内存拷贝

string和[]byte 如何取舍
	string擅长的场景：
		需要字符串比较的场景
		不需要nil字符串的场景
	[]byte擅长的场景:
		修改字符串的场景,尤其是修改粒度为1字节
		函数返回值，需要nil表示含义的场景
		需要切片操作的场景

*/
