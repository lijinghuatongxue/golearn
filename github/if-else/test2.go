package main

import (
	"fmt"
	"runtime"
)

func init()  {
	if runtime.GOOS != "windows"{
		fmt.Println("go的运行环境不是windows")
	}
}


func main()  {
	var key bool = false
//fmt.Println(runtime.GOOS)
	if runtime.GOOS == "darwin"{
		fmt.Println("go的运行环境是mac")
	}  else {
		// return 之后便不能再else了
		fmt.Println(key)
		return
	}
}