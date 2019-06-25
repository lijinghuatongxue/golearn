package main

import (
	"fmt"
	"os"
	"strconv"
)

var key string = "233"
var key2 string = "abc"


func main()  {
	_,err := strconv.Atoi(key)
	if err == nil {
		fmt.Printf("Program stopping with error %v", err)
		//如果这里执行os.exit，程序就会退出
		//os.Exit(1)
	}
	fmt.Printf("上面函数没有写os.exit,所以你看见我了")
	i,err := strconv.Atoi(key)
	if err == nil {
		fmt.Printf("\n 输出err：", err,i+1)
		fmt.Println()
		os.Exit(2)
	}
}