package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
var (
    a  = 'a'
    b  = 'A'
    c  = '!'
    a1 string
    a2 string
    a3 string
)
 
func main() {
    //生成数值
    for i := 1; i <= 26; i++ {
        a1 += string(a)
        a++
        a2 += string(b)
        b++
        a3 += string(c)
        c++
    }
    a4 := a1 + a2 + a3
    bstring := []byte(a4)
    bstringlen := len(bstring)
    var str string
    //输入生成的个数
    var input int
    fmt.Printf("Please input number >>:")
    fmt.Scanln(&input)
    if input > bstringlen {
        fmt.Println("请输入小于", bstringlen)
        // panic("请输入小于", bstringlen)
    } else {
        //生成随机的for
        rand.Seed(time.Now().UnixNano())
        time.Sleep(time.Microsecond)
        for i := 1; i <= input; i++ {
            num := rand.Intn(bstringlen)
            //fmt.Println("NUM=", num)
            s := fmt.Sprintf("%c", bstring[num])
            str += s
 
        }
        fmt.Printf("生产的随机字符=%v\n", str)
 
    }
 
}