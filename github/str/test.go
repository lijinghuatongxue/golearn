package main

import "fmt"


func main()  {
	str1 := "123123sadsadsdsadsaknjknk"

	str2 := "adsdasiiihasdasdsaddsajkn"

	str3 := "asdasda" + "asdasdasdasdsdasdasadasdasdas"
// 拼接
	hello := "hel" + "lo"
	hello += "word"

	fmt.Printf("%d str1长度 \n",len(str1+str2+str3))
	fmt.Println(hello)
}