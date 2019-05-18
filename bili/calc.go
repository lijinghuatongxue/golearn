package main

import "fmt"

// =++++++++++++++++++++++++++++++++++++++++++++ 多个返回值 ++++++++++++++++++++++
// 两个以上的返回值，一定要用括号括起来
func calc(a int , b int)(int,int){
	c := a + b
	d := (a + b)/2
	// 用逗号隔开
	return c ,d
}
func main(){
	sum,avg := calc(100,200)
	fmt.Println("sum=", sum, "avg=", avg)
}