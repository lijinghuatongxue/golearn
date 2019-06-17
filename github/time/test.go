package main

import (
	"fmt"
	"time"
)

func main()  {
	t := time.Now()

	// +++++++++++++++++++++++++++ 未格式化的当前时间 +++++++++++++++++++++++
	fmt.Println("未格式化的当前时间:",time.Now())
	// 打印今天的日期，只打印几日
	fmt.Println(t.Day())
	// 打印今天的日期，打印月份和日期
	fmt.Println(t)
	// ++++++++++++++++++++++++++++ 自己造格式 打印年月日  +++++++++++++++++++
	fmt.Printf("%02d-%02d-%02d\n",t.Year(),t.Month(),t.Day())
	// ++++++++++++++++++++++++++++ 自己造格式 打印年月日小时分钟  +++++++++++++++++++
	fmt.Printf("自己造格式%02d-%02d-%02d-%02d-%02d\n",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute())
	// +++++++++++++++++++++++++++ 自己造格式化的当前时间 +++++++++++++++++++++++
	fmt.Println("自己造格式化的当前时间:",t.Format("2006-06-02-15:04:09"))
	// +++++++++++++++++++++++++++ 纳秒 ++++++++++++++++++++++++
	week := 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)

	// +++++++++++++++++++++++++++ 英文月份转int ++++++++++++++++++++++++
	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)
	// 英文月份转int
	fmt.Println("英文月份转int: ",year, int(month), day)

	// ++++++++++++++++++++++++++ 参考 +++++++++++++++++++++++
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
}