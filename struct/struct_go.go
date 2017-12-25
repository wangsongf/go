package main

import (
    "fmt"
    "reflect"
)

//跟C语言或其它语言一样，也有结构体struct。C语言中用关键词typedef来给结构体定义，Go中用的都是type。struct 语法格式：type typeName struct {...} ,按照对象的理解：struct就是类，函数就是方法。

type person struct {  
    name string  
    age  int32  
}

type course []string

type student struct {  
    name  string  
    age   int32  
    grade string  
}

type class struct {  
    student //匿名字段，struct  
    course  //匿名字段，自定义类型  
    int32   //内置类型做完匿名字段  
    teacher string  
}

type goods struct {  
    name  string  "goods_name"  
    price float64 "goods_price"  
}

func main() {  
    //单独声明和赋值  
    var p1 person  
    p1.name = "zengzhihai"  
    p1.age = 40

    //直接声明和赋值
    p2 := person{"xiaogao", 20}

    //通过 field:value 的方式初始化,这样可以任意顺序
    p3 := person{age: 22, name: "xiaoqing"}

    fmt.Println(p1, p2, p3)

    stu1 := class{student: student{"xiaoming", 18, "高三"}, teacher: "杰伦"}
    //修改course字段
    stu1.course = make([]string, 6)
    stu1.course[0] = "math"
    stu1.course = append(stu1.course, "china")

    //修改匿名内置字段
    stu1.int32 = 5

    fmt.Println(stu1)

    g := &goods{"zhuihui", 40}
    s := reflect.TypeOf(g).Elem() //通过反射获取type的定义
    for i := 0; i < s.NumField(); i++ {
        fmt.Println(s.Field(i).Tag) // 讲tag打印出来
    }
}

