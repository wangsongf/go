//下面是代码例子:
//go 变量
//申明包名main
package main 

//导入包名
import(
    "fmt"
)

var a int //声明一个int类型的变量
var b struct{ //声明一个结构体
    name string
}

var c = 8 //声明变量同时赋值
var ( //批量声明变量，简洁
    d int
    e string
)

var name1 int = 5 //1声明变量name1，并且初始化

//一行申明多个变量
var f,g int

//同一行初始化多个变量,不同类型也可以,这里默认初始化值，根据值进行定义了类型。
var h,i,j = 5,"abd",0.4

func main(){
    name2 := "test" //2声明变量并初始化值，这种赋值不能在函数外面进行赋值
    //打印变量
    fmt.Println(a,b,c,d,e,f,g,h,i,j)
    fmt.Println(name1,name2)
}
