package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	// 将以布尔值的形式去展示
	//fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("%t\n",strings.HasSuffix(str,"ing"))
	fmt.Printf("%t\n",strings.HasSuffix(str,"me"))
}