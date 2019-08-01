package main

// 传递变长参数
import (
	"fmt"
)

func main()  {
	x :=  min(1,3,2,0)
	fmt.Printf("this minimum is: %d\n", x)
	slice := []int{2,5,6,9}
	x = min(slice...)
	fmt.Printf("this mininum in slice is: %d\n", x)
}
func min(s ...int) int{
	if len(s)==0{
		return 0
	}
	min := s[0]
	for _, v := range s{
		if v < min {
			min = v
		}
	}
	return min
}