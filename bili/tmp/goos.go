package main

import (
	"runtime"
	"fmt"
)

func main()  {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n",goos)
}
