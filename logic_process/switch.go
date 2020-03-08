package main

import "fmt"

func main() {

    switch number1 := 3; number1 {
    case 1:
        fmt.Println(1)
    case 3:
        fmt.Println(3)
    case 5:
        fmt.Println(5)
    }
    //不需要写break，一旦条件符合自动结束

    number2 := 6
    switch {
    case number2 == 2:
        fmt.Println(2)
    case number2 == 4:
        fmt.Println(4)
    case number2 == 6:
        fmt.Println(6)
        //如希望继续执行下一个case，可以使用fallthrough语句
        fallthrough
    default:
        fmt.Println("default")
    }

}