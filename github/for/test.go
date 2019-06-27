package main

import "fmt"

func main()  {
	// 标准计算循环
	for i :=  0;i<3;i++{
	fmt.Println("这是一个标准for循环,打印三次")
		for i := 9;i<10;i++{
			fmt.Println("for循环嵌套",i)
			for i := 0;i<2;i++{
				fmt.Println("for循环嵌套",i)
			}
		}
	}
	// 调用字符串
	var str string = "go is good"
	// 空格也算一个长度
	fmt.Println("str的长度是:",len(str))
	for i:=0;i<len(str);i++{
		fmt.Println("打印n" +
			"n次",len(str),str[i])
	}
	//  调用汉字,一个汉字会占用2-4个字节
	var str2 string = "李精华"
	fmt.Println("str2:",len(str2))
	for xx :=1;xx<len(str2);xx++{
		fmt.Printf("Character on position %d is: %c \n",len(str2),str[xx])
	}

}
