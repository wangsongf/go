package main

import (
    "fmt"
    "runtime"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()
        time.Sleep(250 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world") //开一个新的Goroutines执行
    say("hello") //当前Goroutines执行
}