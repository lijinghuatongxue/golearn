package main

//  Harbor 公开仓库检查，null则为正常，目前没有公开的仓库，非null则有
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var url  = "http://root:Swl19960706@47.56.15.133:8888//job/kkk/lastBuild/api/json"
func main() {



	response, err := http.Get(url)
	if err != nil {
		fmt.Println("url error")
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	type Road struct {
		Building   string
		Duration int
	}
	//roads := []Road{
	//	{"Diamond Fork", 29},
	//	{"Sheep Creek", 51},
	//}
	roads := string(body)
	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatalln(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	out.WriteTo(os.Stdout)
}