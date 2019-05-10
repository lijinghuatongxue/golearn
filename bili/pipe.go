package main

import "fmt"

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	var t1 int
	// 变量赋值 需要 =
	t1 = <-pipe
	pipe <- 4
	fmt.Println(t1)
	// python类似，打印长度
	fmt.Println(len(pipe))

}
