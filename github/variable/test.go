package main

import (
	"fmt"
	"os"
	"runtime"
)

// 声明变量的一般形式是使用 var 关键字
// int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。记住，所有的内存在 Go 中都是经过初始化的


// 声明与赋值（初始化）语句也可以组合起来
var a int = 9
var b bool
var str string

var (
	er int = 9
	err bool = false
	stat string
	key = "asd"
)



func main()  {
	fmt.Println(a,err,stat,key)
    // 获取操作系统
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	//这种写法主要用于声明包级别的全局变量，当你在函数体内声明局部变量时，应使用简短声明语法 :=
	//通过 os 包中的函数 os.Getenv() 来获取环境变量中的值
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
