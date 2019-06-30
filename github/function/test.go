package main

import "fmt"

func main()  {
	func1()
	fmt.Println("简单调用函数func1")
}
func func1 ()  {
	fmt.Println("我是函数func1")
}