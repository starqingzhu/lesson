package mutex

import "sync"

/*
mutex

加锁时程序会自动判断是否可以自旋，无限制的自旋将会给CPU带来巨大压力，所以判断是否可以自旋就很重要了。


自旋必须满足一下所有条件。
	1.自旋次数要足够小，通常为4，即自旋最多4次。
	2.cpu核数要大于1，否则自旋没有意义，因为此时不可能有其他协程释放锁。
	3.协程调度机制中的preccess 数量大于1，如果使用GOMAXPROCS（）将处理器设置为1就不能自旋。
	4.协程调度机制中国的可运行队列必须为空，否则会延时协程调度。

自旋的优势：
	更充分利用cpu，尽量避免协程切换。

自旋的问题:
	如果自旋过程中获得锁，那么之前被阻塞的协程将无法获得锁，如果加锁的协程特别多，每次都通过自旋获得锁，那么之前被阻塞的进程将很难获得锁，从而进入饥饿状态。
	为了避免协程长时间无法获取锁，自1.8版本以来增加了一个状态，

Mutex模式：
	normal和starvation模式
	normal模式：
		该模式下，加锁失败协程不会立刻转入阻塞队列中，而是判断是否满足自旋条件，如果满足则会启动自旋过程，尝试抢锁。
	starvation模式：

*/

func MutexFn() {
	var m sync.Mutex
	m.Lock()
	defer m.Unlock()
}

func RWMutexFn() {
	var m sync.RWMutex

	m.RLock()
	defer m.RUnlock()
}
