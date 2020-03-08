/*指针解引用
解引用指针的意思是通过指针访问被指向的值。指针 a 的解引用表示为：*a。
让我们通过一个程序看一下它是怎么工作的。*/
package main
import (
"fmt"
)
func change(val *int) {
	*val = 55
}
func main() {
b := 255
a := &b
fmt.Println("address of b is", a)
fmt.Println("value of b is", *a)
	c := 55
	d := &c
	change(d)
	fmt.Println("value of a after function call is", c)
}