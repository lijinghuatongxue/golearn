package main

import (
	"fmt"
	"github.com/rday/zabbix"
)

func main() {
api, err := zabbix.NewAPI("http://zabbix.geekerland.net/api_jsonrpc.php", "Admin", "***")
if err != nil {
fmt.Println(err)
return
}

versionresult, err := api.Version()
if err != nil {
fmt.Println(err)
}

fmt.Println(versionresult)

_, err = api.Login()
if err != nil {
fmt.Println(err)
return
}
fmt.Println("Connected to API")

}