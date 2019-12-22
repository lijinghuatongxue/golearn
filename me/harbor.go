package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	v := url.Values{}
	v.Set("username", "admin")
	v.Set("password", "")
	//利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}//客户端,被Get,Head以及Post使用
	reqest, err := http.NewRequest("GET", "", body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	////给一个key设定为响应的value.
	//reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value") //必须设定该参数,POST参数才能正常提交

	resp, err := client.Do(reqest)//发送请求
	defer resp.Body.Close()//一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}