package main
// 给出一个n，找出所有a+b=n 模式下的a和b的数字
import (
	"fmt"
)

func list(n int) {

	for i := 0; i <= n; i++ {
		// \t不换行  \n 换行
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}
func main(){
	list(100)
}
