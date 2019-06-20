package main

import "fmt"

// 如果是负数，则把它转正
func Abs(x int) int {
	if x < 0 {
		//return -x
		fmt.Println(-x)
	}
	return x
}
func main()  {
	//var key bool = true
	//if initialization;key{
	//	fmt.Println("222")
	//}
	var a int = 23
	var b int = 45
	Abs(int(-223))

	if a < 50 && b > 10 {
		fmt.Println("a小于50但是大于10")
	}else {
		fmt.Println("不满足要求")
	}
	// 分号在这里起"or"的作用
	if a = 50 ; b < 10 {
		fmt.Println("a小于50但是大于10")
	}else {
		fmt.Println("不满足要求")
	}


}