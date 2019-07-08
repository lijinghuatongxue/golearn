package main

import "fmt"
// return a,b,c 直接返回值
//return
//会在函数内部查找相应的变量，作为返回
// 任何一个非命名返回值（使用非命名返回值是很糟的编程习惯）在 return 语句里面都要明确指出包含返回值的变量或是一个可计算的值

// 非命令返回值
func SumProductdiff(i,j int) ( int, int, int){
	return i+j,i*j,i-j
}
// 命令返回值
func SumProdectDiffN(i,j int)(s int,w int,q int)  {
	s,w,q = i+j,i*j,i-j
	return
}
func main()  {
	sum,prod,diff:=SumProductdiff(7,4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
	sum,prod,diff = SumProdectDiffN(7,3)
	fmt.Println("Sum:",sum,"| Product:",prod,"| Diff:",diff)
}