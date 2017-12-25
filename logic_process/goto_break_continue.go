package main

import "fmt"

func main() {

    for number := 1; number < 5; number++ {
        if number == 3 {
            break
        }
        fmt.Println("break:", number)
    }

    for number := 1; number < 5; number++ {
        if number == 3 {
            continue
        }
        fmt.Println("continue:", number)
    }
lable1:

    for {
        for {
            //配合标签跳出最外层循环
            //标签名是大小写敏感的
            break lable1
        }
    }

    fmt.Println("Hello World1")

    goto lable2

    fmt.Println("Hello World2")

lable2:
    fmt.Println("Hello World3")

}