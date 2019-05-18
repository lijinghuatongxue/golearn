package main
// for循环实现打印1-100
import (
	"fmt"
	"time"
)

func goroute(a int){
	fmt.Println(a)
}
func main(){
	for i := 0; i < 100; i++{
		go goroute(i)
	}
	time.Sleep(time.Second)
}