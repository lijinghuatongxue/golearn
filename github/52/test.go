package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "222"
	var orig2 string = "ABC"
	// var an int
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// 第一个返回值是转换成int的值，第二个返回值判断是否转换成功
	// 这里的是正确的，能转变为数字
	_, err := strconv.Atoi(orig)
	if err != nil{
		fmt.Println("err！！")
	}
	// 这里的是错误的，不能转变为数字，将会打印
	i, err := strconv.Atoi(orig2)
	if err != nil{
		fmt.Println("err！！",i+1)
		}
}