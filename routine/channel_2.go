package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Println(c)
	//这里我们必须从另外一个线程里面写入到通道，然后用主线程或者另外一个线程马上进入等待的状态，这样才能使用无缓冲的channel
	//每一个发送者与接收者都会阻塞当前线程，只有当接受者与发送者都准备就绪了
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}