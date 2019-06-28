package main

import "fmt"

var i int = 5
func main()  {
	for i >=0{
		i--
		fmt.Println(i)
	}
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}
	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
