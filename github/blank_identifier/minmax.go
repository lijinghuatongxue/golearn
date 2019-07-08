package main

import (
	"fmt"
)

func main()  {
	var a int
	var b int
	a,b=minmax(9,6)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", a, b)
}
func minmax(i int,t int)(min int,max int)  {
	if i > t{
		min = i
		max = t
	}
	if  t > i {
		min = i
		max = t
	}
	return
}