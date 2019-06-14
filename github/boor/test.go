package main

import "fmt"

var aVar  bool = true

var bVar bool = false

var cVar bool = false

var key  = false

// ++++++++++++++++++ 和运算符，只有当两边的值都为 true 的时候，和运算符的结果才是 true
// T && T -> true
// T && F -> false
// F && T -> false
// F && F -> false

// +++++++++++++++ 非运算符
// !T -> false
// !F -> true

// 或运算符：||  只有当两边的值都为 false 的时候，或运算符的结果才是 false，其中任意一边的值为 true 就能够使得该表达式的结果为 true
//
// T || T -> true
// T || F -> true
// F || T -> true
// F || F -> false

func main()  {

//	非运算符：!
	fmt.Println("非运算符:",!aVar)
	fmt.Println("和运算符",aVar && bVar,aVar && aVar)

	// 或运算符：||
    fmt.Println(aVar || bVar)
	fmt.Println("或运算符:",cVar || bVar)
	// var 赋值
	fmt.Println("var赋值-非：",!key)
}
