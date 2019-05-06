package main
//引用fmt输出，sync同步
import (
	"fmt"
	"sync"
)
// 构造一个队列
var queue sync.WaitGroup

func printer(num chan int) {
	for i := range num {
		// \n换行
		//fmt.Printf("Received %d \n", i)
		fmt.Printf("%d \n", i)
	}
	//结束队列
	queue.Done()
}
// main是该程序的切入点
// main is the entry point for the program.
func main() {
	//初始化/构造一个channel通道，chan是channel的缩写
	c := make(chan int)
	//开启协程运行，并发运行
	go printer(c)
	//协程,每次+1或者-1
	queue.Add(1)
	// 发送整数到通道
	// Send integers on the channel.
	for i := 1; i <= 10; i++ {
		//把数据放入到协程c
		c <- i
	}
	// 释放内存
	close(c)
	// 等待结束
	queue.Wait()
}