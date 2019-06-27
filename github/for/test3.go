package main

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)
func main()  {
	// 方法1 if判断
	for i :=1;i<101;i++{
		fmt.Println(i)
		if i%FIZZ == 0{
			fmt.Println(i,"3 FIZZ")
		}
		if i%BUZZ == 0{
			fmt.Println(i,"5 BUZZ")
		}
		if i%FIZZBUZZ == 0{
			fmt.Println(i,"15 FIZZBUZZ")
		}
	}

	// 方法2 switch
	for i := 1;i<101;i++{
		switch  {
		case i%FIZZ == 0:
			fmt.Println(i,"3的倍数")
		case i%BUZZ == 0:
			fmt.Println(i,"5的倍数")
		case i%FIZZBUZZ == 0:
			fmt.Println(i,"15的倍数")
		default:
			fmt.Println(i)

		}
	}
}
