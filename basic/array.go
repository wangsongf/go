//go数组和切片
package main 

import(
    "fmt"
)

func main(){  
    //数组定义  
    //数组是Go语言编程中最常用的数据结构之一。顾名思义，数组就是指一系列同一类型数据的集合。数组中包含的每个数据被称为数组元素（ element），一个数组包含的元素个数被称为数组的长度。是值类型。  
    //数组有3种创建方式：[length]Type　、[N]Type{value1, value2, ... , valueN}、[...]Type{value1, value2, ... , valueN}    

    arr1 := [5] int {1,2,3,4,5} //创建数组大小为5
    arr2 := [5] int {1,2} //创建数组大小为5，但是内容没写的用0替代了
    arr3 := [...] int {1,2,3,5} //数组未定义长度
    arr4 := [5] int {2:4,3:5,4:3} //数组有key value 
    arr5 := [...] int {2:3,4:5} //数组长度未定义，并且是key，value形式
    arr6 := [...] string{"my girl","my lift"}
    arr5[1] = 99
    
    //arr5[6] = 88 这是错误的赋值方式，数组不支持最大的key进行赋值。
    fmt.Println(arr1,arr2,arr3,arr4,arr5,arr6)

    arr := [...] int {11,12,13,14,15,16}

    //切片的定义
    //Go语言中，切片是长度可变、容量固定的相同的元素序列。Go语言的切片本质是一个数组。容量固定是因为数组的长度是固定的，切片的容量即隐藏数组的长度。长度可变指的是在数组长度的范围内可变。
    //Go语言提供了数组切片（ slice）这个非常酷的功能来弥补数组的不足。初看起来，数组切片就像一个指向数组的指针，实际上它拥有自己的数据结构，而不仅仅是个指针。数组切片的数据结构可以抽象为以下3个变量： 一个指向原生数组的指针； 数组切片中的元素个数； 数组切片已分配的存储空间。
    //切片的创建有4种方式：1）make ( []Type ,length, capacity ) 2)  make ( []Type, length)  3) []Type{}  4) []Type{value1 , value2 , ... , valueN }

    // s := [] int {1,2,3} 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
    s1 := [] int {1,2,3}

    // s := arr[:] 初始化切片s,是数组arr的引用
    s2 := arr[:]

    // s := arr[startIndex:endIndex] 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
    s3 := arr[0:2]

    // s:= arr[startIndex:] 缺省endIndex时将表示一直到arr的最后一个元素
    s4 := arr[3:]

    // s := arr[:endIndex] 缺省startIndex时将表示从arr的第一个元素开始
    s5 := arr[:3]

    // s := s1[startIndex:endIndex] //通过切片s初始化切片s1
    s6 := s1[1:2]

    // s := make([]int,len,cap) 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
    s7 := make([]int,2)

    //定义string类型的切片
    s8 := []string{"hello", "my boy", "you is beauty"}

    fmt.Println(s1,s2,s3,s4,s5,s6,s7)

    //循环数组的元素
    for i:=0; i<len(arr6); i++ {
        fmt.Println("arr6[",i,"] =", arr6[i])
    }

    //循环数组的元素
    for k,v := range arr6 {
        fmt.Println("arr6[",k,"] =", v)
    }

    //循环切片
    for i:=0;i<len(s8);i++ {
        fmt.Println("s8[",i,"]=",s8[i])
    }
    for k,v := range s8 {
        fmt.Println("s8[",k,"]=",v)
    }

    //切片的追加
    s9 := make([]string,10,20)
    s9 = append(s9,"hehe test")
    fmt.Println(s9)
    s10 := []string{"hello","this","is","new"}
    s10 = append(s8,s9...)
    fmt.Println(s10)
    for k,v := range s10 {
        fmt.Println("s10[",k,"]=",v)
    }
}