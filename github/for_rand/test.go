package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d \n ", a)
	}
	for i := 0; i < 5; i++ {
		// Intn 可以设定最大值
		r := rand.Intn(10)
		fmt.Printf("%d \n ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	fmt.Println(timens)
	//Seed(value) 函数来提供伪随机数的生成种子，一般情况下都会使用当前时间的纳秒级数字
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		// %2.3f 控制小数点前后
		fmt.Printf("%2.3f \n ", 10*rand.Float32())
	}
}