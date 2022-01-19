package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    uri := "https://rules.emergingthreats.net/blockrules/compromised-ips.txt"
    response, err := http.Get(uri)
    if err != nil {
        panic(err)
    } else {
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
    }
}

