package main

import "fmt"

func main() {
	//chan做限速使用
	//不带缓冲的channel写完就阻塞，这种情况只有其他协程中有对应的读才能解除阻塞。而带缓冲长度为N的channel要直到写满+N才阻塞。
	//这里channel的发送超出了缓冲的大小，所以会因为阻塞而导致程序死锁,如果设置channel为3，<-1这种写入操作最多为3次，不然会造成死锁。
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}