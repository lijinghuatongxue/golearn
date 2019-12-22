package main

import "fmt"
// else-if 分支的数量是没有限制的，但是为了代码的可读性，还是不要在 if 后面加入太多的 else-if 结构。如果你必须使用这种形式，则把尽可能先满足的条件放在前面
func main()  {
	var key bool = false
	var key2  = "key2"
	if key{
		fmt.Println("如果成功就打印：",key2)
	}else if key{
		fmt.Println("不成功则打印：",key2)
	}else {
		fmt.Println("上面的一个都没成功")
	}
}

