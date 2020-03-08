
package main

import "fmt"

func main() {
    //使用 make(chan val-type) 创建一个新的通道
    ch := make(chan int)

    //go 关键字创建一个协程
    go loop(ch)
    go loop(ch)

    fmt.Println(ch)

    //使用阻塞的接受方式来等待一个 go 协程的运行结束
    <-ch
    <-ch //从通道中接收值
}

func loop(ch chan int) {
    for i := 0; i < 8; i++ {
        fmt.Println(i)
    }
    ch <- 1 //发送值到通道
}
