package app
// 第三个参数是接受一个管道，int
func Add(a int,b int,c chan int){
	sum := a + b
	// sum的值推给管道c
	c  <- sum
}
