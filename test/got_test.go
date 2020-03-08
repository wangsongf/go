package mytest

import (
	"fmt"
	"testing"
)

//go本身提供一套轻量级的测试框架，符合规则的测试代码会在测试的时候被自动识别并且执行。
//单元测试源文件的命名规则是在需要测试的包下面创建以"_test"结尾的go文件，形式如[^.]*_test.go
//单元测试函数分两类：功能测试函数和性能测试函数，分表以Test和Benchmark为函数名前缀并以*testing.T为单一参数的函数。下面是例子
//这里的测试文件最好单独放到一个文件夹下面，不放到一个文件夹下面。可能会执行其他测试文件
//要执行功能测试 执行：go test mytest.go这样就行
//执行性能测试，执行：go test -bench 6_test.go

func testAdd(t *testing.T) {
	r := 3
	if r != 2 {
		fmt.Print("add err")
	}
}

func BenchmarkAdd2(b *testing.B) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}