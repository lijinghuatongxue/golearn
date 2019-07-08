package main

import "fmt"
//空白符用来匹配一些不需要的值，然后丢弃掉
func main()  {
	var a int
	var c float32
	a,_,c = Threeidentifier()
	fmt.Println("a:",a,"b:","c",c)
}
func Threeidentifier()(int,int,float32)  {
	return 2,3,2.322222
}