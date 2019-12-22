package main

//  Harbor 公开仓库检查，null则为正常，目前没有公开的仓库，非null则有
import (

	"fmt"
	"io/ioutil"
	"net/http"
)

var url  = "https://harbor.rongcloud.net/api/projects?project_name=guest"
func main() {
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("url error")
			defer response.Body.Close()
		}
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}