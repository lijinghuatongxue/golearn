package main

import "fmt"

func main() {
	// label一旦定义，必须要被使用
	// 标签的名称是大小写敏感的，为了提升可读性，一般建议使用全部大写字母
LABEL1:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 3 {
			continue LABEL1
			fmt.Println("label:")
		}
	}
LABEL2:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 3 {
			break  LABEL2
			fmt.Println("label:")
		}
	}
}
