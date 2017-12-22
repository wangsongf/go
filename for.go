
package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i ++ {
		sum += i
	}
	fmt.Println("sum:", sum) // sum: 45

	// 初始化语句和后置语句是可选的
	sum = 1
	for ; sum < 1000; {
		fmt.Println("sum:", sum);
		sum += sum
	}
	fmt.Println("sum:", sum) // sum: 1024

	// 去掉分号 `;`, `for` 是 Go 中的 `while`
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum:", sum) // sum: 1024
	
	//无限循环
	for {
	
    }
}
