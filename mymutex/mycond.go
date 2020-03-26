package mymutex

import ()

/*
条件变量的作用？
	条件变量并不是被用来保护临界区和共享资源的，它是用于协调想要访问共享资源的那些线程的。
	当共享资源的状态发生变化时，它可以被用来通知被互斥锁阻塞的线程

条件变量怎样与互斥锁配合使用？
	条件变量的初始化离不开互斥锁，并且它的方法有的也是基于互斥锁的。
	条件变量提供的方法有三个：等待通知（wait）、单发通知（signal）和广播通知（broadcast）。
*/

func Test() {

	//var mailbox uint8
	//var lock sync.RWMutex
	//sendCond := sync.NewCond(&lock)
	//recvCond := sync.NewCond(lock.RLocker())
}
