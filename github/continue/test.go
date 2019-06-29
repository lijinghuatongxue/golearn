package main

import "fmt"

func main()  {
	for i :=0;i < 100; i++{
		fmt.Println(i)
		if i > 30{
			break
		}
	}
	for i :=0;i < 10; i++{
		fmt.Println(i)
		if i > 3{
			break
		}
	}
}
