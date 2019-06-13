package main

import (
	"fmt"
)
// init执行优先级比 main 函数高
var Pi float64 = 45

func init() {
	//Pi = 4 * math.Atan(1) // init() function computes Pi
	fmt.Println("Pi")
}
func main()  {
	const Pi  = 1
	fmt.Println("main 输出：",Pi)
}