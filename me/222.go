package main

import (
    "bytes"
    "encoding/json"
    "log"
    "os"
)

func main() {

    type Road struct {
        Name   string
        Number int
    }
    roads := []Road{
        {"Diamond Fork", 29},
        {"Sheep Creek", 51},
    }

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
