package main

import (
	"fmt"
)

func main()  {
	var pointer = 22
	// +++++++++++++++++++++++++++++++++ 打印指针的内存地址 +++++++++++++++++++++++
	fmt.Println("指针的地址是：",pointer,&pointer)

	var pointerP *int
	// 当一个指针被定义后没有分配到任何变量时，它的值为 nil
	fmt.Println("定义指针结果：",pointerP)
	pointerP = &pointer
	fmt.Println("指针定义后分配变量结果：",pointerP,"原来的数值为：",pointer)

	key := "chinese"
	var key2 *string = &key
	//通过对 *key2 赋另一个值来更改“对象”，这样 key 也会随之更改
	*key2 = "新值"
	fmt.Println("新内存地址的值：",key2)
	fmt.Println("新值是：",*key2)

}
