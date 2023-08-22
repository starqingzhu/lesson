package memory

/*
内存分配原理：
	为了方便自主管理，做法便是向系统申请一块内存，然后将内存切割成小块，通过一定的内存分配算法管理内存。
	spans+bitmap+arena。
	spans和bitmap是为了管理arena而存在的。
	arena的大小是512G，为了方便管理把arena区域划分成一个个的page，每个page为8KB,一共有512GB/8KB个页。
	spans区域存放的指针，每个指针对应一个page，所以span区域的大小为（512GB/8KB）* 指针大小8byte = 512M
	bitmap区域大小也是通过arena计算出来，不过主要用于GC


span:
	用于管理arena页的关键数据结构，每个span中包含1个或多个连续页，为了满足小对象分配，span中的一页会划分更小的粒度，而对于大对象比如超过页大小，则通过多页实现。

class:
	根据对象的大小，划分了一系列class，每个class都代表一个固定大小的对象，以及每个span的大小。
	结构如下：
	class  bytes/obj	bytes/span	objects		waste bytes
	  1		  8				8192	  1024			0

    class: class ID,每个span结构中都有一个class ID，表示该span可处理的对象类型
	bytes/obj：该class代表对象的字节数
	bytes/span:每个span占用堆的字节数，也即（页数*页大小）
	objects:每个span可分配的对象个数，也即（bytes/span）/(bytes/obj)
	waste bytes: 每个span产生的内存碎片，也即（bytes/spans）%(bytes/obj)



内存逃逸分析
	1.指针逃逸		eg:s :=new(Student)
	2.栈空间不足逃逸
	3.动态类型逃逸    eg: s:="Escape"
	4.闭包引用对象逃逸
		eg: func Fibonacci() func()int {
				a,b := 0,1
				return func() int{
					a,b = b, a+b
					return a
				}
			}


*/
