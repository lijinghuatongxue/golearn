package main

import (
	"fmt"
)

// 打印如下类似
//G
//GG
//GGG
//GGGG
//GGGGG
//GGGGGG


// 方法1 使用 2 层嵌套 for 循环
func main()  {
	for i := 1; i <= 25;i++{
		for g := 1;g < i;g++{
			fmt.Printf("G")
		}
		fmt.Println()
	}

	// 方法2 仅用 1 层 for 循环以及字符串连接
var str string = "G"
for i := 1;i<25;i++{
	fmt.Println(str)
	str += "G"
}

}