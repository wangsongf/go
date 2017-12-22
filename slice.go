
package main

import "fmt"

func main() {
    //`len` 获取`slice`的长度,  `cap` 获取`slice`的最大容量
    number1 := make([]int, 9, 10)
    fmt.Println(len(number1), cap(number1))

    number2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    number3 := number2[:6]
    number4 := number2[6:]
    number5 := number2[4:5]

    fmt.Println(number3, number4, number5)
}
