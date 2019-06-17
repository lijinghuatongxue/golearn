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
	// ++++++++++++++++++++++++++++++++++++  Contains字符串包含关系  +++++++++++++++++++++++++
	var str2 string = "lijinghua is chinese"
	fmt.Printf("\n 包含关系：%s",strings.Contains(str2, "ese"))
	// +++++++++++++++++++++++++++++++++++++ Index字符串索引   ++++++++++++++++++++++++++++
	// 判断子字符串在父字符串的第一个索引位置,一个空格算一个字符
	var child string = "s"
	var father string = "This is an example of a string is is is"
	fmt.Printf("\n 子字符串的第一个索引是: %d",strings.Index(father,child))
	fmt.Printf("\n 子字符串的最后一个索引是：%d",strings.LastIndex(father,child))

	// 非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位
	// strings.IndexRune(s string, r rune) int

	// +++++++++++++++++++++++++++++ Replace字符串替换 +++++++++++++++++++++++++++++++++
	var str3 string = "2019----2019----2019 ---- 2019"
	fmt.Printf("\n 替换结果为：%s",strings.Replace(str3,"2019","9090",1))
	// Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，
	// 如果 n = -1 则替换所有字符串 old 为字符串 new
	//strings.Replace(str, old, new, n) string
	// ++++++++++++++++++++++++++++ Count统计字符出现个数 ++++++++++++++++++++++++++++
	var count string = "10239120382190832122222"
	fmt.Printf("\n 2 的个数是：%d",strings.Count(count,"2"))
	//++++++++++++++++++++++++++++ repeat重复打印字符 ++++++++++++++++++++++++++++
	var repeat string = "2122222222"
	fmt.Printf("\n 重复打印五遍str: %s",strings.Repeat(repeat,5))
	// +++++++++++++++++++++++++++++ ToLower小写 与 ToUpper大写 +++++++++++++++++++++++++++++++
	var xiao string = "uuuGGG"
	fmt.Printf("\n %s",strings.ToLower(xiao))
	fmt.Printf("\n %s",strings.ToUpper(xiao))
	// ++++++++++++++++++++++++++++  Trim 剔除  +++++++++++++++++++++++++++++++++
	var trim string = "     222 	i am chinese  333    "

	fmt.Printf("\n 去除空格的输出：%s",strings.TrimSpace(trim))
	// 这里不能忽视空格
	fmt.Printf("\n 去除左边的特定字符：%s",strings.TrimLeft(trim,"     222"))
	// 去除右边
	fmt.Printf("\n 去除右边的特定字符：%s",strings.TrimRight(trim,"333    "))
	// ++++++++++++++++++++++++++++++ Fields分割字符串 +++++++++++++++++++++++++++++++++++++
	var fields string = "sss aass  sss s a chinese"
	fmt.Printf("\n 分割字符串结果：%s",strings.Fields(fields))
	// ++++++++++++++++++++++++++++++ Split分割字符串 +++++++++++++++++++++++++++++++++++++
	fmt.Printf("\n 按特定字符分割字符串：%s",strings.Split(fields,"sss"))
	// ++++++++++++++++++++++++++++++ 拼接字符串 ++++++++++++++++++++++++++++++++++++++
	var join1 string = "222 222 222"
	var join2 string = "22121"
	fmt.Printf("\n 拼接：%s",strings.Join([]string{join2,join1},"+"))
}