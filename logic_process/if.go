package main

import "fmt"

func main() {

    if number := 7; number < 1 {
        fmt.Println(1)
    } else if number >= 1 && number <= 10 {
        fmt.Println(2)
    } else {
        fmt.Println(3)
    }
    if i := 7; i < 1 {
        fmt.Println(1)
    }

}