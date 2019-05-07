package main

import (
	"fmt"
	"time"
	_"time"
)
func add(a int,b int)int{
	var sum  int
	sum = a + b
	return sum
}
func main(){
	for i := 0; i < 100;i++ {
		go test_print100(i)
	}
	// 跑10s，尽管一秒钟就能跑完，但是程序在2s之后才会退出,继续执行下一个任务
	//time.Sleep(2*time.Second)

	fmt.Println("a+b example ++++++++++++++++++++++++")
	// a+b
	var c int
	c = add(100,200)
	go test_goroute(300,300)
	fmt.Println("a+b =",c)
	fmt.Println("pipe example ++++++++++++++++++++++++")
	// 调用test_pipe example
	test_pipe()

}