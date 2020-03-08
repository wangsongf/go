package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	/*
	结构体字段可以通过结构体指针来访问。

	如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。*/
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	/*结构体文法通过直接列出字段的值来新分配一个结构体。*/
	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)

	fmt.Println(v1, p, v2, v3)
}
