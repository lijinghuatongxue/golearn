package main

import (
	"fmt"
	"io/ioutil"
	//    "log"
	"net/http"
	//    "strings"
	"bytes"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	resp, _ := http.Get("http://10.67.2.252:8080/?a=123456&b=aaa&b=bbb")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var user User
	user.Name = "aaa"
	user.Age = 99
	if bs, err := json.Marshal(user); err == nil {
		//        fmt.Println(string(bs))
		req := bytes.NewBuffer([]byte(bs))
		tmp := `{"name":"junneyang", "age": 88}`
		req = bytes.NewBuffer([]byte(tmp))

		body_type := "application/json;charset=utf-8"
		resp, _ = http.Post("http://10.67.2.252:8080/test/", body_type, req)
		body, _ = ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(err)
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://10.67.2.252:8080/?a=123456&b=aaa&b=bbb", nil)
	request.Header.Set("Connection", "keep-alive")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}

	req := `{"name":"junneyang", "age": 88}`
	req_new := bytes.NewBuffer([]byte(req))
	request, _ = http.NewRequest("POST", "http://10.67.2.252:8080/test/", req_new)
	request.Header.Set("Content-type", "application/json")
	response, _ = client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}