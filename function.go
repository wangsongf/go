package main

import(
    "fmt"
)

/* 定义函数 */
type Circle struct {
      radius float64
}

func main(){  
    //Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：func (variable_name variable_data_type) function_name() [return_type]{ /* 函数体*/}  
    var c1 Circle  
    c1.radius = 10.00  
    fmt.Println("Area of Circle(c1) = ", c1.getArea())

    fmt.Println("---------------")
    //闭包和普通函数的区别
    tmp := []int{1,2,3}
    for _,i := range tmp {
        fmt.Println(i)
        test(i)
    }
    
    fmt.Println("---------------")
    for _,i := range tmp {
        fmt.Println(i)
        //defer延迟关闭改资源，以免引起内存泄漏,defer的执行顺序是逆序的，也就是先进后出的顺序,defer类似析构函数，在函数或者类的最后关闭进行执行。
        defer test(i)
    }

    //此打印结果和上面最好分开打印，不然看不出结果
    fmt.Println("---------------")
    for _,i := range tmp {
        fmt.Println(i)
        //闭包里的非传递参数外部变量值是传引用的，在闭包函数里那个i就是外部非闭包函数自己的参数，所以是相当于引用了外部的变量， i 的值执行到第三次是3 ，闭包是地址引用所以打印了3次i地址指向的值，所以是3，3，3
        defer func() {
            fmt.Println(i)
        }()
    }
    
}

//普通函数
func test(i int){
    fmt.Println(i)
}

//该 method 属于 Circle 类型对象中的方法,计算圆的面子
func (c Circle) getArea() float64 { // func + 主函数 + 函数名(参数) + 返回值类型定义 { 函数内容 }
    //c.radius 即为 Circle 类型对象中的属性
    return 3.14 * c.radius * c.radius
}

