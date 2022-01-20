package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    uri := "https://threatfox-api.abuse.ch/api/v1"
    var json = []byte(` {"query": "get_iocs", "days":7} `)

    // TODO - break this into multiple files: 1 = reusable client and 1 per cti feed
    request, err := http.NewRequest("POST",uri,bytes.NewBuffer(json))
    request.Header.Set("Content-Type","application/json; charset=UTF-8")
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        panic(err)
    } else {
        defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(body))
    }
}
