
package main

import (
    "fmt"
    "math"
)
const Pi = 3.14
func main() {
    fmt.Printf("hello, world\n")
    //变量声明
    var i, j int = 1, 2
    k := 3
    l, m ,n := true,false,"no"
    fmt.Println(i,j,k,l,m,n)
    //变量类型
    var g int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", g, f, b, s)
    //类型转换
    var x, y int = 3, 4
    var u float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(u)
    fmt.Println(x, y, z)
    //常量
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}
