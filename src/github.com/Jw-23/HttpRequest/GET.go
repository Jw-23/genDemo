package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: time.Second * 4}
	req, err := http.NewRequest("GET", "http://localhost:8000", nil)
	if err != nil {
		fmt.Println("newRequest failed,err:", err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Http Get failed,err:", err)
		return
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read response body failed, err:", err)
		return
	}
	fmt.Println(string(bytes))
}
