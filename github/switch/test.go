package main

import "fmt"

func main()  {
	// +++++++++++++++++++++++++++++++++++++++++++++++++++ 形式1 ++++++++++++++++++++++++++++++++++
	var test int = 1
	switch test {
		case 0:
		fmt.Println("输出0")
		case 1:
		fmt.Println("输出1")
		default:
		fmt.Println("默认输出")
	}
	// +++++++++++++++++++++++++++++++++++++++++++++++++++ 形式2 ++++++++++++++++++++++++++++++++++
	var  test2 int = 100
	switch  {
	case test2 == 100:
		fmt.Println("等于100")
	case test2 < 100:
		fmt.Println("小于100")
	default:
		fmt.Println("默认输出")

	}
}
