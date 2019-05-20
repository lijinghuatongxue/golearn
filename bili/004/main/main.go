package main
// 别名前面加一个别名，之后的都能用这个别名
import (
	a "../app"
	"awesomeProject/golearn/bili/004/app"
	"fmt"
)
func main(){
	fmt.Println("Name=",app.Name)
	// 下面这里使用的别名
	fmt.Println("User=",a.User)

}