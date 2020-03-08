package main

import (
	"fmt"
)

func main() {
	//定义一个2维数组，直接赋值
	var arr0 = [5][2]string{{"a", "b"}, {"c", "d"}}
	fmt.Println(arr0)

	//定义一个2维数据根据坐标赋值
	var arr1 = [4][2]string{}
	arr1[0][0] = "aa"
	arr1[0][1] = "bb"
	arr1[1][0] = "cc"
	arr1[1][1] = "dd"
	arr1[2][0] = "ee"
	arr1[2][1] = "ff"
	fmt.Println(arr1)

	//定义一个2维slice，直接赋值
	var slice1 = [][]string{{"aaa", "bbb", "ccc"}, {"ddd", "eee", "fff"}, {"ggg"}}
	fmt.Println(slice1)

	//多维的slice赋值，一定从最里层形成一个slice，然后自里向外逐层append即可。
	//这里是2维slice例子
	var slice2 [][]string
	for i := 0; i < 5; i++ {
		var tmpSlice1 []string
		for j := 0; j < 4; j++ {
			tmpSlice1 = append(tmpSlice1, "a")
		}
		slice2 = append(slice2, tmpSlice1)
	}
	fmt.Println(slice2)
	//这里是3维slice例子
	var slice3 [][][]string
	for i := 0; i < 4; i++ {
		var tmpSlice2 [][]string
		for j := 0; j < 5; j++ {
			var tmpSlice3 []string
			for k := 0; k < 6; k++ {
				tmpSlice3 = append(tmpSlice3, "bb")
			}
			tmpSlice2 = append(tmpSlice2, tmpSlice3)
		}
		slice3 = append(slice3, tmpSlice2)
	}
	fmt.Println(slice3)

}