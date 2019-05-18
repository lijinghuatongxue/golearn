package main

import (
	"fmt"
	"time"
	_"time"
)

//var pipe chan int
func add(a int,b int,c chan int){
	var sum int
	sum = a+b
	time.Sleep(1*time.Second)
	c <- sum
}
func main(){
	// 放在这里则是全局变量了
	var pipe chan int
	pipe = make(chan int,1)
	go add(2,2,pipe)
	sum :=<- pipe
	fmt.Println("sum=",sum)
}
