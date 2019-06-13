package main

// 常量使用关键字 const 定义，用于存储不会改变的数据

import (
	"fmt"
)

// 基础版
const vulue  = 1111


// 多个常量

const (
	a = 1
	b = 2
	c = 3
	d
	e
)



func main(){
	fmt.Println(vulue)
	fmt.Println(a)
	fmt.Println(d)
	// 如果这个常量没有被定义，则默认为最后一个常量的值
	fmt.Println(e)
}