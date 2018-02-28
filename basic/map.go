//go的map和range的使用
package main

import(
    "fmt"
)

func main(){  

    //声明map，map其实就是hashmap的实现  
    var m1 map[string]string  
    //再使用make函数创建一个非nil的map，nil map不能赋值   
    m1 = make(map[string]string)  
    //最后给已申明的map赋值  
    m1["a"] = "aa"  
    m1["b"] = "bb"  

    //直接创建也就是申明加创建
    m2 := make(map[string]string)
    //然后赋值
    m2["c"] = "cc"
    m2["d"] = "dd"

    //初始化 + 声明 + 赋值一体化
    m3 := map[string]string{
        "e": "ee",
        "f": "ff",
    }
    fmt.Println(m1,m2,m3)

    //map 的数据类型：类似于键值对数据结构，语法：map[键数据类型]值数据类型,不是并发安全的
    v_map := map[int]string{
        1:"one",
        2:"two",
        3:"three",
    }
    fmt.Println(v_map);
    fmt.Println(len(v_map))

    //遍历map
    for k,v := range v_map {
        fmt.Println(k,v)
    }
}

