package main

import "fmt"
// 乘法相乘
func main()  {
	fmt.Printf("Multiply的结果是: a*b*c= %d \n",Multiply3Nums(2 , 5 ,6))
}
func Multiply3Nums(a int,b int,c int) int {
	return  a * b * c
}