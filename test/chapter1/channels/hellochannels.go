package main

//引用fmt输出，sync同步
import (
	"fmt"
	"sync"
)
// 构造一个队列
var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		// \n换行
		//fmt.Printf("Received %d \n", i)
		fmt.Printf("%d \n", i)
	}
	wg.Done()
}
// main是该程序的切入点
// main is the entry point for the program.
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)
    // 发送整数到通道
	// Send integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}
	// 释放内存
	close(c)
	// 等待结束
	wg.Wait()
}