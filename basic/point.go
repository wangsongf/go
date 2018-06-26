package main  
import (  
    "fmt"
)

func change(val *int) {  
    *val = 55
}

func main() {  
    b := 255
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
    *a++
    fmt.Println("new value of b is", b)


    c := 58
    fmt.Println("value of c before function call is",c)
    d := &c
    change(d)
    fmt.Println("value of a after function call is", c)
}
