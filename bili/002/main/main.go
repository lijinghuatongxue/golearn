package main

import (
	"../app"
	"fmt"
)
func main(){
	// 声明一个管道
	var pipe chan int
	// 分配内存空间
	pipe = make(chan int,1)
	// 子模块目录名字+函数名字，以goroute形式运行，用go起一个线程，goroute跑完以后，把结果存在pipe(管道)里面
	go app.Add(100,200,pipe)
	sum := <- pipe
	fmt.Println("sum =",sum)
}
