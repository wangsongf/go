package main
import (
	"fmt"
)
func main() {
	/*
	对于go语言来说，设计最精妙的应该是interface了，直白点说interface是一组method的组合。至于更加详细的描述，本文不做介绍，今天谈谈空接口。

	　　空interface(interface{})不包含任何的method，因此所有的类型都实现了空interface。空interface在我们需要存储任意类型的数值的时候相当有用，有点类似于C语言的void*类型。请看下面的代码：
	 */
	slice := make([]interface{}, 10)
	map1 := make(map[string]string)
	map2 := make(map[string]int)
	map2["TaskID"] = 1
	map1["Command"] = "ping"
	map3 := make(map[string]map[string]string)
	map3["mapvalue"] = map1
	slice[0] = map2
	slice[1] = map1
	slice[3] = map3
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[3])
}