package main

import (
	"../calc"
	"fmt"
)
func main(){
	sum := calc.Add(200,100)
	sub := calc.Sub(200,100)
	fmt.Println("sum=",sum)
	fmt.Println("sub=",sub)
}
