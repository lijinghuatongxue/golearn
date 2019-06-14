package main

import "fmt"

var a float64 = 2.234423423
var b float64 = 0.2133232222

var c  float32 = 233.2342
var d  float32 = 234423.23423

var e int64 = 90921312312312312
var f int64 = 233213123

var g int32 = 89789
var j int32 = 1121122

func main()  {
	fmt.Println("float64 a+b=",a+b)
	fmt.Println("float32 c+d=",c+d)
	fmt.Println("int64 e+f=",e+f)
	fmt.Println("int32 g+j=",g+j)
}

