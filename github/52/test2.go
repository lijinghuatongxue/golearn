package main

import (
	"fmt"
	"strconv"
)

var key string = "123"
var key2 string = "abc"
func main()  {
	_,err := strconv.Atoi(key)
	if err == nil{
		fmt.Println("err！！,不能转换int")
		return
		i,err := strconv.Atoi(key2)
		if err != nil{
			fmt.Println("上面没有把err return: ",i+200)
		}
	}

}