package main

// 获取go PATH环境变量
import (
	"fmt"
	"os"
)

func main()  {
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}