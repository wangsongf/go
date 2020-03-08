package main

import (
	"fmt"
	"time"
)

func main() {
	//select语句属于条件分支流程控制语句，不过它只能用于通道。它可以包含若干条case语句，并根据条件选择其中之一执行。select语句的case关键词只能后跟用于通道的发送操作的表达式以及接受操作的表达式或语句。
	//golang 的 select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。

	var ch1 = make(chan int)

	//生成一个协程
	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
	}()
	defer close(ch1)

	done := 0
	finished := 0
	for finished < 3 {
		select {
		case v, ok := <-ch1:
			if ok {
				done = done + 1
				fmt.Println(v)
			}
		}
		finished = finished + 1
	}
	fmt.Println("Done", done)

	//当for 和 select结合使用时，break语言是无法跳出for之外的，因此若要break出来，这里需要加一个标签，使用goto， 或者break 到具体的位置
	//这里是使用break样例
	i := 0
forend:
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("break now")
				break forend
			}
			fmt.Println("inside the select: ")

		}
	}

	//这里使用goto
	i = 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("break now")
				goto ForEnd
			}
			fmt.Println("inside the select: ")
		}
		fmt.Println("inside the for: ")
	}
ForEnd:
}