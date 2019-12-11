package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"


type Toptracks struct {
	Id        string `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string   `json:"lastname"`
	//Address   []Address_info
	Address   Address_info `json:"address"`
}

type Address_info struct {
	Province  string `json:"province"`
	City      string `json:"city"`
}

func get_content() {
	// json data
	url := "http://localhost:12345/people/1"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("%s",body)
	if err != nil {
		panic(err.Error())
	}

	var data Toptracks
//	var data interface{} // TopTracks
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)
	fmt.Println(data.Address.)
	os.Exit(0)
}



func main() {
	get_content()
}
