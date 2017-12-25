package main

import (
    "fmt"
)

var str string //声明字符串

var str1 = "i am a boy"
var str2 = "i am a man"

//字符串函数地址：http://docscn.studygolang.com/pkg/strings/

func main() {
    str = "test"  //赋值
    ch := str[0]  //获取第一个字符
    l := len(str) //字符串长度
    fmt.Println(str, ch, l)

    my := "my name is 曾志海"

    for i := 0; i < len(my); i++ {
        fmt.Println(my[i])
    }

    for i,s := range my {
        fmt.Println(i,"Unicode(",s,") string-",string(s))
    }

    r := []rune(my) //这是一个切片。rune是一种数据类型 rune=int32
    fmt.Println("rune=",r)

    for i:=0;i<len(r);i++{
        fmt.Println("r[",i,"]=",r[i],"string=",string(r[i]))
    }

    fmt.Println(str1,str2)
}

/*
使用说明  
len函数是Go中内置函数，不引入strings包即可使用。len(string)返回的是字符串的字节数。len函数所支持的入参类型如下：  
len(Array) 数组的元素个数  
len(*Array) 数组指针中的元素个数,如果入参为nil则返回0  
len(Slice) 数组切片中元素个数,如果入参为nil则返回0  
len(map) 字典中元素个数,如果入参为nil则返回0  
len(Channel) Channel buffer队列中元素个数 
*/